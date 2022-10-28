package middlewares

import (
	"net/http"

	"vinpel/golang-sample-api-jwt/common/auth"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		headerString := context.GetHeader("Authorization")
		if headerString == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}

		username, err := auth.ValidateToken(headerString, auth.Access)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Set("username", username)
		context.Next()
	}
}
