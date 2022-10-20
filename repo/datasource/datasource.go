package datasource

import (
	"github.com/knullhhf/hack22/models/dto"
	"github.com/knullhhf/hack22/models/entity"
	"github.com/knullhhf/hack22/repo"
)

var DataSourceImplement DataSourceRepoInterface = &DataSourceRepo{}

type DataSourceRepoInterface interface {
	List() ([]*entity.DataSource, error)
	ListDatabases(datasourceId int) ([]*dto.DatabaseDto, error)
	ListTables(database string, datasourceId int) ([]*dto.TableDto, error)
	Find(id int) ([]*entity.DataSource, error)
}

type DataSourceRepo struct {
}

func (DataSourceRepo DataSourceRepo) List() ([]*entity.DataSource, error) {
	datasources := []*entity.DataSource{}
	err := repo.GetDB().Model(datasources).Debug().Find(&datasources).Error
	return datasources, err
}

func (DataSourceRepo DataSourceRepo) Find(id int) ([]*entity.DataSource, error) {
	datasources := []*entity.DataSource{}
	err := repo.GetDB().Model(datasources).Debug().Where("id = ?", id).Find(&datasources).Error
	return datasources, err
}

func (DataSourceRepo DataSourceRepo) ListDatabases(datasourceId int) ([]*dto.DatabaseDto, error) {
	ds, err := DataSourceRepo.Find(datasourceId)
	if err != nil {
		return nil, err
	}
	db, err := repo.OpenDb(ds[0])
	if err != nil {
		return nil, err
	}
	var databases []*dto.DatabaseDto
	err = db.Raw("select distinct(table_schema) as name from information_schema.tables ").Scan(&databases).Error
	return databases, err
}

func (DataSourceRepo DataSourceRepo) ListTables(database string, datasourceId int) ([]*dto.TableDto, error) {
	ds, err := DataSourceRepo.Find(datasourceId)
	if err != nil {
		return nil, err
	}
	db, err := repo.OpenDb(ds[0])
	if err != nil {
		return nil, err
	}

	var tables []*dto.TableDto
	err = db.Raw("select table_name as name from information_schema.tables where table_schema=? ", database).Scan(&tables).Error
	return tables, err
}
