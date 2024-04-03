package controllers

import (
	"os"
	"log"
	"fmt"
	"net/http"

	"github.com/andyfusniak/base58"
	"github.com/gin-gonic/gin"

	"url-shortner/internal/initializers"
	"url-shortner/internal/models"
)

type ShortenUrlPayload struct {
	URL string `json:"url"`
}

func ShortenUrl(c *gin.Context) {
	log.Println("Shorten url called")

	var shortUrl models.ShortUrl

	err := c.BindJSON(&shortUrl)

	if err != nil {
		// Return a 400 Bad Request response if JSON binding fails
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create a slug
	slug, err := base58.RandString(8)

	if err != nil {
		// Log the error and return a 500 Internal Server Error response
		log.Println("Error creating shorturl:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	shortUrl.Slug = slug

	result := initializers.DB.Create(&shortUrl)
	if result.Error != nil {
		// Log the error and return a 500 Internal Server Error response
		log.Println("Error creating shorturl:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// shorturl entry created
	log.Println("Short url entry created")

	// building url
	var scheme = "https"
	if os.Getenv("USE_TLS") == "False" {
		scheme = "http"
	}

	var domain = os.Getenv("DOMAIN")

	c.IndentedJSON(
		http.StatusCreated,
		gin.H{
			"msg": "Created",
			"url": fmt.Sprintf("%s://%s/%s", scheme, domain, slug),
		},
	)
}
