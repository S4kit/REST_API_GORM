package database

import (
	"log"
	"os"

	"github.com/sixfwa/fiber-api/Models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type dbInstance struct {
	Db *gorm.DB
}

var Database dbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the Database", err.Error())
		os.Exit(2)
	}
	log.Println("Connected to the Database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migration")
	// TODO : Add Migrations
	db.AutoMigrate(&Models.User{}, &Models.Product{}, &Models.Order{})
	Database = dbInstance{Db: db}
}
