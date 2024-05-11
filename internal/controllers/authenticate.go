package controllers

import (
	"log"
	"net/http"

	"github.com/YogeshUpdhyay/url-shortner/internal/initializers"
	"github.com/YogeshUpdhyay/url-shortner/internal/models"
	"github.com/YogeshUpdhyay/url-shortner/internal/serializers"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Authenticate(ctx *gin.Context) {
	// parsing the payload from the context
	request := serializers.AuthenticateRequest{}

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		log.Println("Authenticate: There was error binding data to request obj")
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"msg":    "Bad Request",
				"detail": err.Error(),
			},
		)
		return
	}

	// validating the request body
	err = request.Validate()
	if err != nil {
		log.Println("Authenticate: There was error validating data the request")
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"msg":    "Bad Request",
				"detail": err.Error(),
			},
		)
		return
	}

	// run business logic
	user := models.User{}
	result := initializers.DB.First(&user, "email = ?", request.Email)
	if result.Error != nil {
		log.Println("Unable to find the user")
		if result.Error == gorm.ErrRecordNotFound {
			log.Println("Unable to find the user")
			ctx.AbortWithStatusJSON(
				http.StatusNotFound,
				gin.H{
					"msg":    "Not Found",
					"detail": "User not found",
				},
			)
		}
	}

	// compare password hashes to verify the identity
	// generate a user token
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		log.Println("Incorrect password")
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"msg":    "Unauthorized",
				"detail": "Incorrect password.",
			},
		)
		return
	}

}
