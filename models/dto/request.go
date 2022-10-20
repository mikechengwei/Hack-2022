package dto

type ListTablesRequestDto struct {
	DataSourceId int    `json:"datasourceId"`
	Database     string `json:"database"`
}

type ListDatabaseRequestDto struct {
	DataSourceId int `json:"datasourceId"`
}
