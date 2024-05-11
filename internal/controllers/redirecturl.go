package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/YogeshUpdhyay/url-shortner/internal/initializers"
	"github.com/YogeshUpdhyay/url-shortner/internal/models"
)

func RedirectUrl(c *gin.Context) {
	log.Println("Redirect url called", c.Param("slug"))
	slug := c.Param("slug")

	// query the shorturls db to get the url
	var shortUrl models.ShortUrl
	result := initializers.DB.First(&shortUrl, "slug = ?", slug)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Url not found"})
		return
	}

	c.Status(301)
	c.Header("Location", shortUrl.URL)
	c.Abort()
}
