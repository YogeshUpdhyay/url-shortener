package controllers

import (
	"log"

	"net/http"
	"github.com/gin-gonic/gin"
)

func RedirectUrl(c *gin.Context) {
	log.Println("Redirect url called", c.Param("urlId"))
	c.JSON(http.StatusOK, gin.H{"msg": "OK"})
}
