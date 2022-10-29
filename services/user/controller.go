package user

import (
	"vinpel/golang-sample-api-jwt/common/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(apiGroup *gin.RouterGroup) {

	userRoutes := apiGroup.Group("/user")
	{
		userRoutes.POST("/register", RegisterUser)
		userRoutes.POST("/change-password", ChangeUserPassword)
		userRoutes.Use(middlewares.Auth())
		{
			userRoutes.GET("/", GetUser)
			userRoutes.GET("/:id", GetUserById)
		}
	}
}
