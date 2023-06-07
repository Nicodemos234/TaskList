package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	EndDate   time.Time      `json:"endDate"`
	Daily     bool           `json:"daily"`
	Done      bool           `json:"done"`
	Title     string         `json:"title"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type CreateTaskInput struct {
	EndDate time.Time `json:"endDate"`
	Daily   bool      `json:"daily"`
	Title   string    `json:"title"`
	Done    bool      `json:"done"`
}

type UpdateTaskInput struct {
	EndDate time.Time `json:"endDate"`
	Daily   bool      `json:"daily"`
	Title   string    `json:"title"`
	Done    bool      `json:"done"`
}
