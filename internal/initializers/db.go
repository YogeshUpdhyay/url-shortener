package initializers

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"url-shortner/internal/models"
)

var DB *gorm.DB

func InitializeDatabase() {
	var err error

	dbName := os.Getenv("DB_NAME")
	DB, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})

	if err != nil {
		log.Fatal("There was an error connnecting to the database")
	}

	// add migration here
	DB.AutoMigrate(&models.ShortUrl{})
}
