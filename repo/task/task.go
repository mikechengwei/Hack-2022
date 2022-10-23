package task

import (
	"github.com/knullhhf/hack22/models/entity"
	"github.com/knullhhf/hack22/repo"
	"github.com/sirupsen/logrus"
)

var TaskRepoImplement TaskRepoInterface = &TaskRepo{}

type TaskRepoInterface interface {
	CreateTask(task *entity.Task) error

	ListTask(pageNumber int, name string) ([]*entity.Task, *int64, error)

	GetTask(taskId int) (result *entity.Task, err error)

	UpdateStatus(taskId int, status int) error
}
type TaskRepo struct {
}

func (t *TaskRepo) CreateTask(task *entity.Task) error {
	r := repo.GetDB().Create(&task)
	return r.Error
}

func (t *TaskRepo) ListTask(pageNumber int, name string) ([]*entity.Task, *int64, error) {
	var results []*entity.Task
	r := repo.GetDB().Limit(10).Where("name LIKE ? ", "%"+name+"%").Order("updated_at DESC").Offset((pageNumber - 1) * 10).Find(&results).Error

	var total int64
	err := repo.GetDB().Model(&results).Where("name LIKE ? ", "%"+name+"%").Count(&total).Error
	if err != nil {
		logrus.Errorf("查询用户应用表失败 err:%v", err)
		return nil, nil, err
	}
	return results, &total, r
}

func (t *TaskRepo) GetTask(taskId int) (result *entity.Task, err error) {
	err = repo.GetDB().Where("id = ?", taskId).Find(&result).Error
	return result, err
}

func (ac *TaskRepo) UpdateStatus(taskId int, status int) error {
	task := &entity.Task{}
	if err := repo.GetDB().Debug().Table(task.TableName()).Where("id = ?", taskId).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}
