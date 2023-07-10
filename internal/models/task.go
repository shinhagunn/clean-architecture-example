package models

import "time"

type TaskStatus string

var (
	TaskStatusProcessing = TaskStatus("processing")
	TaskStatusDone       = TaskStatus("done")
)

type Task struct {
	ID         int64      `gorm:"type:bigint;not null;autoIncrement"`
	UserID     int64      `gorm:"type:bigint;notnull"`
	CategoryID int64      `gorm:"type:bigint;notnull"`
	Level      int64      `gorm:"type:bigint;notnull"`
	Name       string     `gorm:"type:character varying(30);not null"`
	Status     TaskStatus `gorm:"type:character varying(10);not null;default:processing"`
	DeadlineAt time.Time  `gorm:"type:timestamp;not null"`
	CreatedAt  time.Time  `gorm:"type:timestamp;not null"`
	UpdatedAt  time.Time  `gorm:"type:timestamp;not null"`
}

func (t Task) TableName() string {
	return "tasks"
}
