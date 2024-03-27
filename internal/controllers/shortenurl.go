package controllers

import (
	"log"

	"net/http"
	"github.com/gin-gonic/gin"
)

func ShortenUrl(c *gin.Context) {
	log.Println("Shorten url called")
	c.JSON(http.StatusOK, gin.H{"msg": "OK"})
}
