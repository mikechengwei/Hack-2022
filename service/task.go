package service

import (
	"context"
	"fmt"
	"github.com/knullhhf/hack22/common"
	"github.com/knullhhf/hack22/logger"
	"github.com/knullhhf/hack22/models/dto"
	"github.com/knullhhf/hack22/models/entity"
	task3 "github.com/knullhhf/hack22/pkg/models/task"
	"github.com/knullhhf/hack22/pkg/server"
	"github.com/knullhhf/hack22/repo/datasource"
	task2 "github.com/knullhhf/hack22/repo/task"
	"github.com/pingcap/tidb/br/pkg/lightning/config"
	"github.com/pingcap/tidb/br/pkg/lightning/log"
	"strconv"
	"strings"
	"time"
)

var TaskServiceImplement TaskServiceInterface = &TaskService{}

type TaskServiceInterface interface {
	CreateTask(task *dto.CreateTaskRequestDto) error
	ListTask(task *dto.ListTaskRequestDto) (*dto.TaskResponse, error)
	StartTask(taskId int) error
	GetTaskProgress(taskId int) (string, error)
}

type TaskService struct {
}

func (t *TaskService) CreateTask(task *dto.CreateTaskRequestDto) error {
	var isSyncSchema int
	if task.IsSyncSchema {
		isSyncSchema = 1
	} else {
		isSyncSchema = 0
	}
	item := &entity.Task{
		Name:             task.Name,
		SourceClient:     task.Source.Client,
		SourceDatabase:   task.Source.Database,
		TargetDatabase:   task.Target.Database,
		SourceDatasource: task.Source.Datasource,
		SourceTables:     strings.Join(task.Source.Table, ","),
		SourceSql:        task.Source.SelectSql,
		SourceSplitMode:  task.Source.TaskSplitMode,
		TargetDatasource: task.Target.Datasource,
		TargetImportMode: task.Target.ImportMode,
		Concurrent:       task.Concurrent,
		SyncSchema:       isSyncSchema,
		FinishedAt:       time.Now(),
		Status:           0,
	}
	err := task2.TaskRepoImplement.CreateTask(item)
	if err != nil {
		return common.CommonError{Code: common.DbError, Detail: err}
	}
	return nil
}

func (t *TaskService) ListTask(task *dto.ListTaskRequestDto) (*dto.TaskResponse, error) {
	tasks, total, err := task2.TaskRepoImplement.ListTask(task.PageNumber, task.Name)
	if err != nil {
		return nil, common.CommonError{Code: common.DbError, Detail: err}
	}

	var taskDto []*dto.TaskDto
	for _, t := range tasks {
		dto := &dto.TaskDto{
			Name:       t.Name,
			CreateTime: t.CreatedAt,
			FinishTime: t.FinishedAt,
			Client:     t.SourceClient,
			Status:     t.Status,
			ID:         t.ID,
		}
		taskDto = append(taskDto, dto)
	}
	resp := &dto.TaskResponse{
		Items:      taskDto,
		Total:      *total,
		PageNumber: task.PageNumber,
	}
	return resp, nil

}

func (t *TaskService) StartTask(taskId int) error {
	task, err := task2.TaskRepoImplement.GetTask(taskId)
	if err != nil {
		logger.LogErrf("get task err:%v", err)
	}
	targetDataSource, err := datasource.DataSourceImplement.Find(task.TargetDatasource)
	if err != nil || len(targetDataSource) < 1 {
		logger.LogErrf("get targetDataSource err:%v", err)
	}

	sourceDataSource, err := datasource.DataSourceImplement.Find(task.SourceDatasource)
	if err != nil || len(sourceDataSource) < 1 {
		logger.LogErrf("get sourceDataSource err:%v", err)
	}
	cli := task.SourceClient
	cli = server.LightningServer.KeyWithNameMap[cli]
	for {
		time.Sleep(time.Second)
		_, e := server.LightningServer.FindCli(cli)
		if e == nil {
			break
		}
		logger.LogInfof("find client[%s] success", cli)
	}

	ctx := context.TODO()
	cfg := config.NewConfig()
	cfg.LoadFromGlobal(&config.GlobalConfig{
		App: config.GlobalLightning{
			Config: log.Config{
				Level: "debug",
			},
		},
	})
	cfg.Adjust(ctx)

	cfg.Checkpoint = config.Checkpoint{
		Schema: "tidb_lightning_checkpoint",
		DSN:    "/tmp/tidb_lightning_checkpoint.pb",
		Driver: "file",
		Enable: true,
	}

	port, _ := strconv.Atoi(targetDataSource[0].Port)
	statusPort, _ := strconv.Atoi(targetDataSource[0].StatusPort)
	cfg.App.TableConcurrency = 1
	cfg.TikvImporter.RangeConcurrency = 16
	cfg.App.IndexConcurrency = 2
	cfg.Mydumper.CSV.Header = true
	cfg.Mydumper.CSV.Header = true
	cfg.App.CheckRequirements = false
	cfg.TiDB.Host = targetDataSource[0].Host
	cfg.TiDB.Port = port
	cfg.TiDB.User = targetDataSource[0].Username
	cfg.TiDB.StatusPort = statusPort
	cfg.TiDB.PdAddr = targetDataSource[0].PdAddress
	cfg.TiDB.Psw = targetDataSource[0].Password
	cfg.Mydumper.Filter = []string{"*.*", "!mysql.*", "!sys.*", "!INFORMATION_SCHEMA.*", "!PERFORMANCE_SCHEMA.*", "!METRICS_SCHEMA.*", "!INSPECTION_SCHEMA.*"}
	cfg.TikvImporter.Backend = "local"
	cfg.TikvImporter.SortedKVDir = "/tmp/lightning"

	sourcePort, _ := strconv.Atoi(sourceDataSource[0].Port)

	tables := strings.Split(task.SourceTables, ",")
	targetDatabase := task.SourceDatabase
	if len(task.TargetDatabase) > 0 {
		targetDatabase = task.TargetDatabase
	}
	for _, table := range tables {
		migrateTask := &task3.MigrateTask{
			ClientName: cli,
			Name:       task.Name,
			Key:        task.Name,
			TaskId:     task.ID,
			Source: &task3.TableInfo{
				Host:     sourceDataSource[0].Host,
				Port:     int32(sourcePort),
				Username: sourceDataSource[0].Username,
				Password: sourceDataSource[0].Password,
				Type:     sourceDataSource[0].Type,
				Database: task.SourceDatabase, Name: table,
				ServiceName: sourceDataSource[0].ServiceName,
			},
			Target: &task3.TableInfo{
				Database: targetDatabase, Name: table,
			},
			Config: cfg,
		}

		server.LightningServer.AddTask(migrateTask)
	}
	return nil

}

func (t *TaskService) GetTaskProgress(taskId int) (string, error) {
	task, err := task2.TaskRepoImplement.GetTask(taskId)
	if err != nil {
		return "", common.CommonError{Code: common.DbError, Detail: err}
	}
	database := task.TargetDatabase
	tables := strings.Split(task.SourceTables, ",")
	var result []string
	for _, table := range tables {
		distotal, total, err := datasource.DataSourceImplement.CountTargetDbData(task.TargetDatasource, database, table)
		if err != nil {
			continue
		}
		result = append(result, fmt.Sprintf("源表数据: %d|%s:%s 完成搬运: %d", distotal, database, table, total))
	}

	return strings.Join(result, "\n"), nil

}
