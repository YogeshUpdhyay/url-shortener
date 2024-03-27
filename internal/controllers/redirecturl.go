package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"url-shortner/internal/initializers"
	"url-shortner/internal/models"
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

	// redirecting the url
	c.Redirect(http.StatusTemporaryRedirect, shortUrl.URL)
}
