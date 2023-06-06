package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	EndDate   time.Time      `json:"endDate"`
	Daily     bool           `json:"daily"`
	Title     string         `json:"title"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type CreateTaskInput struct {
	EndDate time.Time `json:"endDate"`
	Daily   bool      `json:"daily"`
	Title   string    `json:"title"`
}

type UpdateTaskInput struct {
	EndDate time.Time `json:"endDate"`
	Daily   bool      `json:"daily"`
	Title   string    `json:"title"`
}
