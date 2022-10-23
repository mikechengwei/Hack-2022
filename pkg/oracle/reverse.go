package oracle

import (
	"context"
	"fmt"
	"strings"

	"github.com/wentaojin/transferdb/pkg/reverse"
	"github.com/wentaojin/transferdb/pkg/reverse/o2m"
	"github.com/wentaojin/transferdb/service"
	"github.com/wentaojin/transferdb/utils"
)

type Reverser struct {
	oraConf   *OracleConfig
	mysqlConf *MySQLConfig

	ctx    context.Context
	engine *service.Engine
}

func NewReverser(ctx context.Context, oraConf *OracleConfig, mysqlConf *MySQLConfig) *Reverser {
	oradb, err := NewOracleDBEngine(oraConf)
	if err != nil {
		panic(err)
	}

	if err := NewMySQLDBPrepareEngine(mysqlConf); err != nil {
		panic(err)
	}
	gormDB, mysqlDB, err := NewMySQLDBRunEngine(mysqlConf)
	if err != nil {
		panic(err)
	}
	dbHandle := &service.Engine{
		OracleDB: oradb,
		MysqlDB:  mysqlDB,
		GormDB:   gormDB,
	}

	if err := dbHandle.InitMysqlEngineDB(); err != nil {
		panic(err)
	}
	if err := dbHandle.InitDefaultValueMap(); err != nil {
		panic(err)
	}
	return &Reverser{
		oraConf:   oraConf,
		mysqlConf: mysqlConf,
		ctx:       ctx,
		engine:    dbHandle,
	}
}

func (r *Reverser) Reverse() error {
	sourceSchema := r.oraConf.SchemaName

	// oracle db collation
	nlsSort, err := r.engine.GetOracleDBCharacterNLSSortCollation()
	if err != nil {
		return err
	}
	nlsComp, err := r.engine.GetOracleDBCharacterNLSCompCollation()
	if err != nil {
		return err
	}
	if _, ok := utils.OracleCollationMap[strings.ToUpper(nlsSort)]; !ok {
		return fmt.Errorf("oracle db nls sort [%s] , mysql db isn't support", nlsSort)
	}
	if _, ok := utils.OracleCollationMap[strings.ToUpper(nlsComp)]; !ok {
		return fmt.Errorf("oracle db nls comp [%s] , mysql db isn't support", nlsComp)
	}
	if !strings.EqualFold(nlsSort, nlsComp) {
		return fmt.Errorf("oracle db nls_sort [%s] and nls_comp [%s] isn't different, need be equal; because mysql db isn't support", nlsSort, nlsComp)
	}
	// reverse 表任务列表
	var tbls []o2m.Table

	// oracle 环境信息
	characterSet, err := r.engine.GetOracleDBCharacterSet()
	if err != nil {
		return err
	}
	if _, ok := utils.OracleDBCharacterSetMap[strings.Split(characterSet, ".")[1]]; !ok {
		return fmt.Errorf("oracle db character set [%v] isn't support", characterSet)
	}

	// oracle 版本是否可指定表、字段 collation
	// oracle db nls_sort/nls_comp 值需要相等，USING_NLS_COMP 值取 nls_comp
	oraDBVersion, err := r.engine.GetOracleDBVersion()
	if err != nil {
		return err
	}

	oraCollation := false
	if utils.VersionOrdinal(oraDBVersion) >= utils.VersionOrdinal(utils.OracleTableColumnCollationDBVersion) {
		oraCollation = true
	}

	var (
		tblCollation    map[string]string
		schemaCollation string
	)

	if oraCollation {
		schemaCollation, err = r.engine.GetOracleSchemaCollation(sourceSchema)
		if err != nil {
			return err
		}
		tblCollation, err = r.engine.GetOracleTableCollation(sourceSchema, schemaCollation)
		if err != nil {
			return err
		}
	}

	tablesMap, err := r.engine.GetOracleTableType(sourceSchema)
	if err != nil {
		return err
	}

	// 库名、表名规则
	tbl := o2m.Table{
		SourceSchemaName: strings.ToUpper(sourceSchema),
		TargetSchemaName: strings.ToUpper(r.mysqlConf.SchemaName),
		SourceTableName:  strings.ToUpper(r.oraConf.Table),
		TargetDBType:     "TiDB",
		TargetTableName:  strings.ToUpper(r.oraConf.Table),
		SourceTableType:  tablesMap[strings.ToUpper(r.oraConf.Table)],
		SourceDBNLSSort:  nlsSort,
		SourceDBNLSComp:  nlsComp,
		Overwrite:        false,
		Engine:           r.engine,
	}
	tbl.OracleCollation = oraCollation
	if oraCollation {
		tbl.SourceSchemaCollation = schemaCollation
		tbl.SourceTableCollation = tblCollation[strings.ToUpper(r.oraConf.Table)]
	}

	tbls = append(tbls, tbl)

	for _, t := range tbls {
		// 表名转换
		modifyTableName := reverse.ChangeOracleTableName(t.SourceTableName, t.TargetTableName)
		t.TargetTableName = modifyTableName

		tableStruct, _, err := t.GenCreateTableSQL(modifyTableName)
		if err != nil {
			return err
		}

		fmt.Println(tableStruct)
		// 直写 tidb
		if _, err := t.Engine.MysqlDB.ExecContext(r.ctx, tableStruct); err != nil {
			return err
		}
	}
	return nil
}
