package utils

import (
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

	privateKey, err := jwt.ParseECPrivateKeyFromPEM(privateKeyBytes)
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

type CustomClaims struct {
	jwt.RegisteredClaims
	UserId int `json:"userId"`
}

func ValidateUserToken(tokenString string) (int, error) {
	signingKeyBytes, err := os.ReadFile("signingkey.pem")
	if err != nil {
		log.Println("ValidateUserToken: Error reading the signing key file")
		return 0, err
	}

	privateKey, err := jwt.ParseECPrivateKeyFromPEM(signingKeyBytes)
	if err != nil {
		log.Println(err.Error())
		log.Println("ValidateUserToken: Error reading the publickey data")
		return 0, err
	}

	publicKey := &privateKey.PublicKey

	// decoding using the public key
	userClaims := CustomClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &userClaims, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if err != nil {
		log.Println(err.Error())
		log.Println("ValidateUserToken: Error validating the token.")
		return 0, err
	}

	// fetching claims from the token
	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		log.Println("ValidateUserToken: Error fetching claims from the token.")
		return 0, constants.ErrFetchingClaimsFromToken
	}

	return claims.UserId, nil
}
