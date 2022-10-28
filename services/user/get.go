package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"vinpel/golang-sample-api-jwt/common/database"
	"vinpel/golang-sample-api-jwt/models"
)

// @Summary     GetUser
// @Schemes     http
// @Description Create a user account
// @Tags        User
// @Produce     json
// @Success     201    {object}    []models.User "Created : return data for the new user"
// @Failure     400       {object} object{error=string}
// @Failure     500       {object} object{error=string}
// @Router      /api/user/  [get]
// @Security    JWTToken
// @Param       Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func GetUser(context *gin.Context) {
	var user []models.User

	db := database.Instance
	db.Find(&user)

	context.JSON(http.StatusOK, user)

}

// @Summary     GetUserById
// @Schemes     http
// @Description Create a user account
// @Tags        User
// @Produce     json
// @Success     201    {object}    []models.User "Created : return data for the new user"
// @Param       id  path    int true "id "
// @Failure     400       {object} object{error=string}
// @Failure     500       {object} object{error=string}
// @Router      /api/user/{id}  [get]
// @Security    JWTToken
// @Param       Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func GetUserById(context *gin.Context) {
	var user []models.User
	id := context.Param("id")
	db := database.Instance
	db.Where("id= ?", id).Find(&user)

	context.JSON(http.StatusOK, user)

}
