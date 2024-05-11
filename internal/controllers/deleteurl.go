package controllers

import (
	"log"

	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/YogeshUpdhyay/url-shortner/internal/initializers"
	"github.com/YogeshUpdhyay/url-shortner/internal/models"
	"github.com/YogeshUpdhyay/url-shortner/internal/serializers"
)

func DeleteUrl(c *gin.Context) {
	log.Println("Delete the url")

	// parsing the payload
	var payload serializers.DeleteUrlRequest
	err := c.BindJSON(&payload)
	if err != nil {
		// Return a 400 Bad Request response if JSON binding fails
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var shortUrl models.ShortUrl
	result := initializers.DB.First(&shortUrl, "slug = ?", payload.Slug)
	if result.Error != nil {
		// Log the error and return a 500 Internal Server Error response
		log.Println("Error deleting shorturl:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	deleteResult := initializers.DB.Delete(&shortUrl)
	if deleteResult.Error != nil {
		// Log the error and return a 500 Internal Server Error response
		log.Println("Error deleting shorturl:", deleteResult.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "OK"})
}
