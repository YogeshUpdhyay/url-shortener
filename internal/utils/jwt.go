package utils

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"log"
	"os"

	"github.com/YogeshUpdhyay/url-shortner/internal/constants"
	"github.com/golang-jwt/jwt/v5"
)

func CreateUserAccessToken(userId string) (string, error) {
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

	token := jwt.New(jwt.SigningMethodES256)
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		log.Println("CreateUserAccessToken: Error creaing the token string")
		return constants.Empty, err
	}

	return tokenString, nil
}
