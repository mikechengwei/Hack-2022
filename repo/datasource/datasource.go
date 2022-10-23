package datasource

import (
	"fmt"
	"github.com/knullhhf/hack22/models/dto"
	"github.com/knullhhf/hack22/models/entity"
	"github.com/knullhhf/hack22/repo"
	"github.com/sirupsen/logrus"
)

var DataSourceImplement DataSourceRepoInterface = &DataSourceRepo{}

type DataSourceRepoInterface interface {
	List() ([]*entity.DataSource, error)
	ListDatabases(datasourceId int) ([]*dto.DatabaseDto, error)
	ListTables(database string, datasourceId int) ([]*dto.TableDto, error)
	Find(id int) ([]*entity.DataSource, error)
	CountTargetDbData(datasourceId int, database string, table string) (int, int, error)
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
	if ds[0].Type == 1 {
		return []*dto.DatabaseDto{
			{Name: "hackathon"},
		}, nil
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

	if ds[0].Type == 1 {
		return []*dto.TableDto{
			{Name: "marvin"},
		}, nil
	}
	db, err := repo.OpenDb(ds[0])
	if err != nil {
		return nil, err
	}

	var tables []*dto.TableDto
	err = db.Raw("select table_name as name from information_schema.tables where table_schema=? ", database).Scan(&tables).Error
	return tables, err
}

func (DataSourceRepo DataSourceRepo) CountTargetDbData(datasourceId int, database string, table string) (int, int, error) {
	ds, err := DataSourceRepo.Find(datasourceId)
	if err != nil {
		return 0, 0, err
	}
	db, err := repo.OpenDb(ds[0])
	if err != nil {
		return 0, 0, err
	}
	total := &dto.TaskTotal{}
	err = db.Raw(fmt.Sprintf("select count(*) as total from %s.%s ", database, table)).Scan(total).Error
	if err != nil {
		logrus.Errorf("查询用户应用表失败 err:%v", err)
		return 0, 0, err
	}
	return total.Total, total.Total, err
}
