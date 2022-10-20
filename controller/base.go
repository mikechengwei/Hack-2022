package controller

import (
	"errors"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/knullhhf/hack22/common"
	"github.com/knullhhf/hack22/models/dto"
)

type BaseController struct {
	beego.Controller
}

func (b *BaseController) ErrorResp(data interface{}, code common.Code, err error) {
	if common.IsInternalCode(code) {
		code = common.APISystemError
	}
	resp := &dto.Response{
		Data:    data,
		Code:    code,
		Message: err.Error(),
	}
	b.Data["json"] = resp
	b.ServeJSON()
	b.StopRun()
}

func ErrorCode(err error) common.Code {
	var e common.CommonError
	if errors.As(err, &e) {
		return e.Code
	}
	return common.SYSTEMERROR
}

func (b *BaseController) Error(data interface{}, err error) {
	code := ErrorCode(err)
	if common.IsInternalCode(code) {
		code = common.APISystemError
	}
	resp := &dto.Response{
		Data:    data,
		Code:    code,
		Message: err.Error(),
	}
	b.Data["json"] = resp
	b.ServeJSON()
	b.StopRun()
}

func (b *BaseController) SuccessResp(data interface{}) {
	resp := dto.Response{
		Code:    common.APISuccess,
		Message: "请求成功",
		Data:    data,
	}
	b.Data["json"] = resp
	b.ServeJSON()
	b.StopRun()
}
