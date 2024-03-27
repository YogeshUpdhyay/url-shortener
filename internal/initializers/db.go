package initializers

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDatabase() {
	var err error
	DB, err := gorm.Open(sqlite.Open("urlshortener.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("There was an error connnecting to the database")
	}

	// add migration here
	DB.AutoMigrate()
}
