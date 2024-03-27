package routes

import (
	"url-shortner/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/shortenurl", controllers.ShortenUrl)
	router.GET("/:urlId", controllers.RedirectUrl)
	router.POST("/deleteurl", controllers.DeleteUrl)
}
