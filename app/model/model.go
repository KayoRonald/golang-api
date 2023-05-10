package model

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	BookID      string     `gorm:"type:uuid;primaryKey" json:"user_id"`
	Title       string     `json:"title"`
	UrlImage    string     `json:"urlImagem"`
	Description string     `json:"description"`
	Price       uint       `json:"price"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}
