package internal

import (
	"github.com/gin-gonic/gin"

	"url-shortner/internal/initializers"
	"url-shortner/internal/routes"
)

var router *gin.Engine

func GetApp() *gin.Engine {
	router := gin.Default()

	// init database
	initializers.InitializeDatabase()

	// registering routes
	routes.RegisterRoutes(router)

	return router
}
