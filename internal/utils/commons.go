package utils

import (
	"log"

	"github.com/YogeshUpdhyay/url-shortner/internal/constants"
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
