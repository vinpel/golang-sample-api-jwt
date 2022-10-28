package user

import (
	"net/http"

	"vinpel/golang-sample-api-jwt/common/database"
	"vinpel/golang-sample-api-jwt/common/dto"
	"vinpel/golang-sample-api-jwt/common/rand"
	"vinpel/golang-sample-api-jwt/models"

	"github.com/gin-gonic/gin"
)

// @Summary     RegisterUser
// @Schemes     http
// @Description Create a user account
// @Tags        User
// @Produce     json
// @Success     201       {object} object{email=string,user_id=string,username=string} "Created : return data for the new user"
// @Failure     400       {object} object{error=string}
// @Failure     500       {object} object{error=string}
// @Param       bodyParam body     dto.UserRequest true "Required data to create a user"
// @Router      /api/user/register  [post]
func RegisterUser(context *gin.Context) {
	var user models.User
	var userRequest dto.UserRequest

	if err := context.ShouldBindJSON(&userRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	user.CreateFromQuery(&userRequest)

	if err := user.CreatePassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	user.Empreinte = rand.String(12)
	record := database.Instance.Create(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{"user_id": user.ID, "email": user.Email, "username": user.Username})
}
