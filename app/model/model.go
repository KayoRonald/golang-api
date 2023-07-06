package model

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	BookID      string     `gorm:"type:uuid;primaryKey" json:"book_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}
