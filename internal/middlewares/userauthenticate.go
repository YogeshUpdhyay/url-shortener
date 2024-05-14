package middlewares

import "github.com/gin-gonic/gin"

func UserAuthenticateMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
