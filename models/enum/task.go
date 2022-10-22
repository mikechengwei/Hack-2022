package enum

type DataSourceType int32

const (
	Oracle DataSourceType = 1
	Mysql  DataSourceType = 2
	TiDB   DataSourceType = 3
)

var DataSourceMap = map[int32]DataSourceType{
	1: Oracle,
	2: Mysql,
	3: TiDB,
}
