package service

import (
	"github.com/knullhhf/hack22/common"
	"github.com/knullhhf/hack22/models/dto"
	"github.com/knullhhf/hack22/models/entity"
	task2 "github.com/knullhhf/hack22/repo/task"
	"strings"
)

var TaskServiceImplement TaskServiceInterface = &TaskService{}

type TaskServiceInterface interface {
	CreateTask(task *dto.CreateTaskRequestDto) error
	ListTask(task *dto.ListTaskRequestDto) (*dto.TaskResponse, error)
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
