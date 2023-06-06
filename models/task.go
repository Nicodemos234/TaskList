package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	EndDate   time.Time
	Daily     bool
	Title     string
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type CreateTaskInput struct {
	Title   string
	EndDate time.Time
	Daily   bool
}

type UpdateTaskInput struct {
	Title   string
	EndDate time.Time
	Daily   bool
}
