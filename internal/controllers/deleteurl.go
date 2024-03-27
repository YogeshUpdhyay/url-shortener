package controllers

import (
	"log"

	"net/http"
	"github.com/gin-gonic/gin"
)

func DeleteUrl(c *gin.Context) {
	log.Println("Deleten the url")
	c.JSON(http.StatusOK, gin.H{"msg": "OK"})
}
