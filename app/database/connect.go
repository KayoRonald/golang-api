package database

import (
	"fmt"

	"github.com/KayoRonald/golang-api/app/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	DB, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(
		&model.Book{},
	)
	fmt.Println("Database Migrated")
}