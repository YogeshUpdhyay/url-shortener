package routes

import (
	"github.com/gin-gonic/gin"

	"url-shortner/internal/controllers"
)

func RegisterRoutes(router *gin.Engine) {
	apiRouter := router.Group("/api/v1")
	{
		apiRouter.POST("/shortenurl", controllers.ShortenUrl)
		apiRouter.POST("/deleteurl", controllers.DeleteUrl)
	}
	router.GET("/:slug", controllers.RedirectUrl)
	router.GET("/", controllers.RedirectToApp)
}