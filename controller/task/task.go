package task

import (
	"encoding/json"
	"github.com/knullhhf/hack22/common"
	"github.com/knullhhf/hack22/controller"
	"github.com/knullhhf/hack22/models/dto"
	"github.com/knullhhf/hack22/service"
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
	case "listTasks":
		t.ListTask()
	default:
		t.ErrorResp("action不支持", common.APICodeNotFoundPath, common.Newf("动作不支持"))
	}
}

func (t *TaskController) ListExportMode() {
	t.SuccessResp([]dto.TaskExportMode{
		{Name: "Logical Import Mode", Id: 1},
		{Name: "Physical Import Mode", Id: 2},
	})
}

func (t *TaskController) ListTaskMode() {
	t.SuccessResp([]dto.TaskMode{
		{Name: "按表拆分"},
		{Id: 1},
	})
}

func (t *TaskController) CreateTask() {
	var request *dto.CreateTaskRequestDto
	if err := json.Unmarshal(t.Ctx.Input.RequestBody, &request); err != nil {
		t.ErrorResp(nil, common.APIParameterError, err)
		return
	}
	err := service.TaskServiceImplement.CreateTask(request)
	if err != nil {
		t.Error(nil, err)
	}
	t.SuccessResp(nil)
}

func (t *TaskController) ListTask() {

	var request *dto.ListTaskRequestDto
	if err := json.Unmarshal(t.Ctx.Input.RequestBody, &request); err != nil {
		t.ErrorResp(nil, common.APIParameterError, err)
		return
	}
	result, err := service.TaskServiceImplement.ListTask(request)
	if err != nil {
		t.Error(nil, err)
	}
	t.SuccessResp(result)
}

func (t *TaskController) StartTask() {
}
