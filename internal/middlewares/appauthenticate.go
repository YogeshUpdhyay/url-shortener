package middlewares

import (
	"log"
	"net/http"

	"github.com/YogeshUpdhyay/url-shortener/internal/constants"
	"github.com/YogeshUpdhyay/url-shortener/internal/initializers"
	"github.com/YogeshUpdhyay/url-shortener/internal/models"
	"github.com/gin-gonic/gin"
)

// authenticate apps that are using the url-shortener as service
func AppAuthenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get the api key header that is used to authenticate the app request
		// validating the api key
		apiKey := ctx.GetHeader(constants.ApiKeyHeaderKey)

		if apiKey == constants.Empty {
			log.Println("Api Key not found")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "Missing api key header."})
			return
		}

		// based on the api key set the appname in context for tracking the request
		credential := models.Credential{}
		result := initializers.DB.First(&credential, "apiKey = ?", apiKey)
		if result.Error != nil {
			log.Println("There was an error fetching the app context based on the api key")
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"msg":    "Invalid api key",
					"detail": "No app found for the api key",
				},
			)
			return
		}

		// setting app name in context
		ctx.Set(constants.ContextAppNameKey, credential.AppName)
		ctx.Next()
	}
}
