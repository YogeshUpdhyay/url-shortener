package controllers

import (
	"log"
	"net/http"

	"github.com/YogeshUpdhyay/url-shortener/internal/serializers"
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
	// apiKey, err
	// credential := models.Credential{
	// 	AppName: request.AppName,
	// 	ApiKey: ,
	// }
}
