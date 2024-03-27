package internal

import (
	"url-shortner/internal/routes"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func GetApp() *gin.Engine {
	router := gin.Default()

	// registering routes
	routes.RegisterRoutes(router)

	return router
}
