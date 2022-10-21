package dto

type ListTablesRequestDto struct {
	DataSourceId int    `json:"datasourceId"`
	Database     string `json:"database"`
}

type ListDatabaseRequestDto struct {
	DataSourceId int `json:"datasourceId"`
}

type CreateTaskRequestDto struct {
	Name         string `json:"name"`
	Source       Source `json:"source"`
	Target       Target `json:"target"`
	Concurrent   int    `json:"concurrent"`
	IsSyncSchema bool   `json:"isSyncSchema"`
}

type ListTaskRequestDto struct {
	PageNumber int    `json:"pageNumber"`
	Name       string `json:"name"`
}
type StartTaskRequestDto struct {
	TaskId int `json:"taskId"`
}
type Source struct {
	Client        string   `json:"client"`
	Datasource    int      `json:"datasource"`
	Database      string   `json:"database"`
	Table         []string `json:"table"`
	SelectSql     string   `json:"selectSql"`
	TaskSplitMode int      `json:"taskSplitMode"`
}

type Target struct {
	Datasource int    `json:"datasource"`
	Database   string `json:"database"`
	ImportMode int    `json:"importMode"`
}
