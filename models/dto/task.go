package dto

import (
	"time"
)

type TaskExportMode struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type TaskMode struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type TaskResponse struct {
	PageNumber int        `json:"pageNumber"`
	Total      int64      `json:"total"`
	Items      []*TaskDto `json:"items"`
}

type TaskDto struct {
	ID         uint      `json:"id"`
	Name       string    `json:"name"`
	Client     string    `json:"client"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"createTime"`
	FinishTime time.Time `json:"finishTime"`
}

type TaskProgressDto struct {
	Status  int    `json:"status"`
	Content string `json:"content"`
}

type TaskTotal struct {
	Total int `json:"total"`
}
