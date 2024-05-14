package utils

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"log"
	"os"
	"time"

	"github.com/YogeshUpdhyay/url-shortener/internal/constants"
	"github.com/golang-jwt/jwt/v5"
)

func CreateUserAccessToken(userId uint) (string, error) {
	privateKeyBytes, err := os.ReadFile("signingkey.pem")
	if err != nil {
		log.Println("CreateUserAccessToken: Error reading the signing key file")
		return constants.Empty, err
	}

	privateKeyData := bytes.NewReader(privateKeyBytes)
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), privateKeyData)
	if err != nil {
		log.Println("CreateUserAccessToken: Error reading the privatekey data")
		return constants.Empty, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * constants.TokenExpirationHours).Unix(),
	})
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		log.Println("CreateUserAccessToken: Error creaing the token string")
		return constants.Empty, err
	}

	return tokenString, nil
}

func ValidateUserToken(tokenString string) error {
	return nil
}
