package initializers

import (
	"log"
	"os"

	"github.com/YogeshUpdhyay/url-shortener/internal/constants"
	"github.com/YogeshUpdhyay/url-shortener/internal/models"
	"github.com/YogeshUpdhyay/url-shortener/internal/utils"
)

func InitSystemDefaults() {
	// validate root user creds
	rootUserEmail := os.Getenv(constants.RootUserEmailEnvKey)
	rootUserPassword := os.Getenv(constants.RootUserPasswordEnvKey)
	if rootUserEmail == constants.Empty || rootUserPassword == constants.Empty {
		log.Fatal("InitSystemDefaults: Error missing system dfault credentials.")
	}

	// add root user with admin permission to the system
	hasehedPassword, err := utils.GetHashedPasswordString(rootUserPassword)
	if err != nil {
		log.Fatal("InitSystemDefaults: Error initializing the system defaults")
	}

	user := models.User{
		Email:    rootUserEmail,
		Password: hasehedPassword,
	}

	result := DB.Create(&user)
	if result.Error != nil && result.Error.Error() != "UNIQUE constraint failed: users.email" {
		log.Fatal("InitSystemDefaults: Error initializing the system defaults")
	}

	log.Println("InitSystemDefaults: System default initialized successfully.")
}
