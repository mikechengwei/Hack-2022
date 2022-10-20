package dto

type DataSourceDto struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type DatabaseDto struct {
	Name string `json:"name"`
}

type TableDto struct {
	Name string `json:"name"`
}
