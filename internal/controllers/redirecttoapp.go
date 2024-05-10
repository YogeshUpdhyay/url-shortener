package controllers

import (
	"os"
	"log"
	"fmt"

	"github.com/gin-gonic/gin"
)

func RedirectToApp(c *gin.Context) {
	applicationUrl := fmt.Sprintf("http://%s", os.Getenv("DOMAIN"))
	log.Println(applicationUrl)
	c.Status(302)
	c.Header("Location", applicationUrl)
	c.Abort()
}
