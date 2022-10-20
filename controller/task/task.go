package task

import (
	"github.com/knullhhf/hack22/common"
	"github.com/knullhhf/hack22/controller"
	"github.com/knullhhf/hack22/models/dto"
)

type TaskController struct {
	controller.BaseController
}

func (t *TaskController) Post() {
	switch t.Ctx.Request.Header.Get("action") {
	case "listExportMode":
		t.ListExportMode()
	case "listTaskMode":
		t.ListTaskMode()
	case "createTask":
		t.CreateTask()
	default:
		t.ErrorResp("action不支持", common.APICodeNotFoundPath, common.Newf("动作不支持"))
	}
}

func (t *TaskController) ListExportMode() {
	t.SuccessResp([]dto.TaskExportMode{
		{Name: "Logical Import Mode"},
		{Id: 1},
		{Name: "Physical Import Mode"},
		{Id: 2},
	})
}

func (t *TaskController) ListTaskMode() {
	t.SuccessResp([]dto.TaskMode{
		{Name: "按表拆分"},
		{Id: 1},
	})
}

func (t *TaskController) CreateTask() {
}

func (t *TaskController) ListTask() {
}

func (t *TaskController) StartTask() {
}
