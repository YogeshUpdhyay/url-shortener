package utils

import (
	"crypto/rand"
	"encoding/base64"
	"log"

	"github.com/YogeshUpdhyay/url-shortener/internal/constants"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetHeaders(ctx *gin.Context) map[string]string {
	headers := map[string]string{}
	// process all headers and return in this dict
	for key, value := range ctx.Request.Header {
		headers[key] = value[0]
	}
	return headers
}

func GetHashedPasswordString(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("GenerateHashedPassword: Error generating password hash")
		return constants.Empty, err
	}

	return string(hashedPassword), nil
}

func GenerateRandomString() (string, error) {
	// Create a byte slice to store the random bytes
	randomBytes := make([]byte, constants.ApiKeyLength)

	// Read random bytes from crypto/rand into the byte slice
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// Encode the random bytes to base64 to get a random string
	randomString := base64.StdEncoding.EncodeToString(randomBytes)

	return randomString, nil
}
