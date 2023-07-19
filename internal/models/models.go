package models

import "time"

type Model struct {
	ID        int64     `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"type:timestamp;not null"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null"`
}
