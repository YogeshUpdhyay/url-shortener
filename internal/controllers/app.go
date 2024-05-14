package controllers

import (
	"log"
	"net/http"

	"github.com/YogeshUpdhyay/url-shortener/internal/initializers"
	"github.com/YogeshUpdhyay/url-shortener/internal/models"
	"github.com/YogeshUpdhyay/url-shortener/internal/serializers"
	"github.com/YogeshUpdhyay/url-shortener/internal/utils"
	"github.com/gin-gonic/gin"
)

func CreateApp(ctx *gin.Context) {
	// parse the request
	request := serializers.CreateAppRequest{}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		log.Println("CreateApp: Error binding request to object")
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"msg":    "Bad Request",
				"detail": err.Error(),
			},
		)
		return
	}

	// validate the request
	err = request.Validate()
	if err != nil {
		log.Println("CreateApp: Error validating the request body")
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
	apiKey, err := utils.GenerateRandomString()
	if err != nil {
		log.Println("CreateApp: Error generating the apikey")
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"msg":    "InternalServer",
				"detail": err.Error(),
			},
		)
		return
	}

	credential := models.Credential{
		AppName: request.AppName,
		ApiKey:  apiKey,
	}
	result := initializers.DB.Create(&credential)
	if result.Error != nil {
		log.Println("CreateApp: Error creating app credntials.")
		log.Println(result.Error.Error())
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"msg":    "Internal Server Error.",
				"detail": result.Error.Error(),
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, &serializers.CreateAppResponse{
		AppName: credential.AppName,
		ApiKey:  credential.ApiKey,
		ID:      credential.ID,
	})
}

func HandleListApps(ctx *gin.Context) {
	// querying all applications and pagination
	apps := []models.Credential{}

	result := initializers.DB.Find(&apps)
	if result.Error != nil {
		log.Println("HandleListApps: Error querying the app credentials.")
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"msg":    "Internal Server Error",
				"detail": "Error querying the app credentials.",
			},
		)
		return
	}

}
