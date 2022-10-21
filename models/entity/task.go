package entity

import "time"

type Task struct {
	BaseEntity
	Name             string
	SourceClient     string
	SourceDatasource int
	SourceTables     string
	SourceSql        string
	SourceSplitMode  int
	SourceDatabase   string
	TargetDatabase   string
	Status           int
	TargetDatasource int
	TargetImportMode int
	Concurrent       int
	SyncSchema       int
	FinishedAt       time.Time
}

func (task *Task) TableName() string {
	return "hackathon_task"
}
