package models

import "time"

type Category struct {
	ID        int64     `gorm:"type:bigint;not null;autoIncrement"`
	Name      string    `gorm:"type:character varying;not null;unique"`
	CreatedAt time.Time `gorm:"type:timestamp;not null"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null"`
}

func (c Category) TableName() string {
	return "categories"
}
