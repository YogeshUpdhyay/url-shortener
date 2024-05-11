package internal

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/YogeshUpdhyay/url-shortner/internal/initializers"
	"github.com/YogeshUpdhyay/url-shortner/internal/routes"
)

var router *gin.Engine

func GetApp() *gin.Engine {
	// init env
	initializers.InitializeEnv()

	router = gin.Default()

	// CORS middleware configuration
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type"}

	// Use CORS middleware
	router.Use(cors.New(config))

	// init database
	initializers.InitializeDatabase()

	// init system defaults
	initializers.InitSystemDefaults()

	// registering routes
	routes.RegisterRoutes(router)

	return router
}
