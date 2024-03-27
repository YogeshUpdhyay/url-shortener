package routes

import (
	"github.com/gin-gonic/gin"

	"url-shortner/internal/controllers"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/shortenurl", controllers.ShortenUrl)
	router.GET("/:slug", controllers.RedirectUrl)
	router.POST("/deleteurl", controllers.DeleteUrl)
}
