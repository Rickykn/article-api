package models

import (
	"time"

	"gorm.io/gorm"
)

type Articel struct {
	ID        int    `json:"id" gorm:"primary_key"`
	Author    string `json:"author"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
