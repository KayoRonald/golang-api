package database

import (
	"os"

	"github.com/KayoRonald/golang-api/app/configs"
	"github.com/KayoRonald/golang-api/app/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	// "gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var (
	log *configs.Logger
	Database DbInstance
)

// var 

func ConnectDB() {
	log = configs.GetLogger("database")
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

	if err != nil {
		log.Error("Failed to connect to the database! \n", err)
		os.Exit(2)
	}
	log.Info("Connected Successfully to Database")
	
	// db.Logger = logger.Default.LogMode(logger.Info)
	log.Info("Running Migrations")
	db.AutoMigrate(&models.Book{})

	Database = DbInstance{
		Db: db,
	}
}
