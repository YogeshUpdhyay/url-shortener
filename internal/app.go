package internal

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/YogeshUpdhyay/url-shortener/internal/initializers"
	"github.com/YogeshUpdhyay/url-shortener/internal/routes"
)

var router *gin.Engine

func GetApp() *gin.Engine {
	// init env
	initializers.InitializeEnv()

	router = gin.Default()

	// CORS configuration
	config := cors.Config{
		AllowAllOrigins:  true,                                                // Allow all origins
		AllowMethods:     []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"}, // Allowed HTTP methods
		AllowHeaders:     []string{"*"},                                       // Allow all headers
		ExposeHeaders:    []string{"Content-Length"},                          // Optional: expose additional headers
		AllowCredentials: true,                                                // Optional: allow credentials if needed
		MaxAge:           0,                                                   // Disable caching of the preflight response
	}

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
