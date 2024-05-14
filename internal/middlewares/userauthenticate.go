package middlewares

import (
	"net/http"
	"strings"

	"github.com/YogeshUpdhyay/url-shortener/internal/constants"
	"github.com/YogeshUpdhyay/url-shortener/internal/utils"
	"github.com/gin-gonic/gin"
)

func UserAuthenticateMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get the auth header a validate
		authHeader := ctx.GetHeader(constants.AuthHeaderKey)

		// Check if the Authorization header is missing
		if authHeader == "" {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"msg":    "Unauthorized",
					"detail": "Authorization header is missing",
				},
			)
			return
		}

		// check if the Authorization header has the Bearer prefix
		if !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"msg":    "Unauthorized",
					"detail": "Invalid authorization format",
				},
			)
			return
		}

		// extract the token from the Authorization header
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// perform token validation (you can implement your own validation logic here)
		userId, err := utils.ValidateUserToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"msg":    "Unauthorized",
					"detail": "Invalid token.",
				},
			)
			return
		}

		// setting user id for the context
		ctx.Set(constants.AuthenticatedUserContextKey, userId)

		ctx.Next()
	}
}
