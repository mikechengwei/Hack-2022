package datasource

import (
	"encoding/json"
	"github.com/knullhhf/hack22/common"
	"github.com/knullhhf/hack22/controller"
	"github.com/knullhhf/hack22/models/dto"
	"github.com/knullhhf/hack22/service"
)

type DataSourceController struct {
	controller.BaseController
}

func (d *DataSourceController) Post() {
	switch d.Ctx.Request.Header.Get("action") {
	case "ListDataSource":
		d.ListDataSource()
	case "ListDataBase":
		d.ListDataBase()
	case "listTables":
		d.ListTables()
	default:
		d.ErrorResp("action不支持", common.APICodeNotFoundPath, common.Newf("动作不支持"))
	}
}

func (d *DataSourceController) ListDataSource() {
	result, err := service.DataSourceServiceImplement.List()
	if err != nil {
		d.Error(nil, err)
	}
	d.SuccessResp(result)
}

func (d *DataSourceController) ListDataBase() {
	var request dto.ListDatabaseRequestDto
	if err := json.Unmarshal(d.Ctx.Input.RequestBody, &request); err != nil {
		d.ErrorResp(nil, common.APIParameterError, err)
		return
	}
	result, err := service.DataSourceServiceImplement.ListDatabases(request.DataSourceId)
	if err != nil {
		d.Error(nil, err)
		return
	}
	d.SuccessResp(result)
}

func (d *DataSourceController) ListTables() {
	var request dto.ListTablesRequestDto
	if err := json.Unmarshal(d.Ctx.Input.RequestBody, &request); err != nil {
		d.ErrorResp(nil, common.APIParameterError, err)
		return
	}
	result, err := service.DataSourceServiceImplement.ListTables(request.DataSourceId, request.Database)
	if err != nil {
		d.Error(nil, err)
	}
	d.SuccessResp(result)
}
