package service

import (
	"github.com/knullhhf/hack22/common"
	"github.com/knullhhf/hack22/models/dto"
	"github.com/knullhhf/hack22/pkg/utils/mapper"
	"github.com/knullhhf/hack22/repo/datasource"
)

var DataSourceServiceImplement DataSourceServiceInterface = &DataSourceService{}

type DataSourceServiceInterface interface {
	List() ([]*dto.DataSourceDto, error)
	ListDatabases(datasourceId int) ([]*dto.DatabaseDto, error)
	ListTables(datasourceId int, database string) ([]*dto.TableDto, error)
}

type DataSourceService struct {
}

func (ds *DataSourceService) List() ([]*dto.DataSourceDto, error) {
	datasources, err := datasource.DataSourceImplement.List()
	if err != nil {
		return nil, common.CommonError{Code: common.DbError, Detail: err}
	}
	var datasourceDtos []*dto.DataSourceDto
	mapper.Copy(&datasourceDtos, datasources)
	return datasourceDtos, nil
}

func (ds *DataSourceService) ListDatabases(datasourceId int) ([]*dto.DatabaseDto, error) {
	databases, err := datasource.DataSourceImplement.ListDatabases(datasourceId)
	if err != nil {
		return nil, common.CommonError{Code: common.DbError, Detail: err}
	}
	var databasesDtos []*dto.DatabaseDto
	mapper.Copy(&databasesDtos, &databases)
	return databasesDtos, nil
}

func (ds *DataSourceService) ListTables(datasourceId int, database string) ([]*dto.TableDto, error) {
	tables, err := datasource.DataSourceImplement.ListTables(database, datasourceId)
	if err != nil {
		return nil, common.CommonError{Code: common.DbError, Detail: err}
	}
	var databasesDtos []*dto.TableDto
	mapper.Copy(&databasesDtos, &tables)
	return tables, nil
}
