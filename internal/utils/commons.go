package utils

import "github.com/gin-gonic/gin"

func GetHeaders(ctx *gin.Context) map[string]string {
	headers := map[string]string{}
	// process all headers and return in this dict
	for key, value := range ctx.Request.Header {
		headers[key] = value[0]
	}
	return headers
}
