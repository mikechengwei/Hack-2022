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

type TaskState int32

const (
	Pending  DataSourceType = 0
	Running  DataSourceType = 1
	Failed   DataSourceType = 2
	Finished DataSourceType = 3
)

var TaskStateMap = map[int32]DataSourceType{
	0: Pending,
	1: Running,
	2: Failed,
	3: Finished,
}
