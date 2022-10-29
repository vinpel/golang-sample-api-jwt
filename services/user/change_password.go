package user

import (
	"net/http"

	"vinpel/golang-sample-api-jwt/common/database"
	"vinpel/golang-sample-api-jwt/common/rand"
	"vinpel/golang-sample-api-jwt/models"

	"github.com/gin-gonic/gin"
)

type changePasswordRequest struct {
	UserName          string `json:"username"  example:"john"`
	ActualPassword    string `json:"actual-password"  example:"demo"`
	NewPassword       string `json:"new-password"  example:"demo2"`
	LogoutAllSessions bool   `json:"logout-all-sesssion"  example:"true"`
}

// @Summary     ChangeUserPassword
// @Schemes     http
// @Description Change the password of the user
// @Tags        User
// @Produce     json
// @Success     200       {object} object{sucess=string,message=string}
// @Failure     400       {object} object{error=string}
// @Failure     500       {object} object{error=string}
// @Param       bodyParam body    changePasswordRequest true "new and old password"
// @Router      /api/user/change-password  [post]
func ChangeUserPassword(context *gin.Context) {
	var requestData changePasswordRequest
	var user models.User

	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	result := database.Instance.Where("username = ?", requestData.UserName).First(&user)
	if result.RowsAffected != 1 {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "cont' find the unique username"})
		context.Abort()
		return
	}

	if err := user.VerifyPassword(requestData.ActualPassword); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	if err := user.CreatePassword(requestData.NewPassword); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	if requestData.LogoutAllSessions {
		user.Empreinte = rand.String(12)
	}

	record := database.Instance.Save(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"user_id": user.ID, "username": user.Username})
}
