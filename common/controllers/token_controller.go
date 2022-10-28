package controllers

import (
	"log"
	"net/http"

	"vinpel/golang-sample-api-jwt/common/auth"
	"vinpel/golang-sample-api-jwt/common/database"
	"vinpel/golang-sample-api-jwt/common/dto"
	"vinpel/golang-sample-api-jwt/models"

	"github.com/gin-gonic/gin"
)

// @Summary     GenerateTokenFromUser
// @Schemes     http
// @Description Generate a new  access and refresh token from a refresh token
// @Description Refresh token is requested  for this endpoint
// @Description  * Refresh token is valid 15 days
// @Description  * Access token is valid 1 hour
// @Tags        Jwt token
// @Produce     json
// @Success     200       {object} auth.JWTTokens
// @Failure     401       {object} object{error=string}
// @Failure     500       {object} object{error=string}
// @Param       bodyParam body     dto.TokenRequest true "Identifiants"
// @Router      /api/token/from-user  [post]
func GenerateTokenFromUser(context *gin.Context) {
	var request dto.TokenRequest
	var user models.User
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	// check if email exists and password is correct
	record := database.Instance.Where("email = ?", request.Email).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	log.Println(request.Password)
	credentialError := user.VerifyPassword(request.Password)
	if credentialError != nil {
		log.Println(credentialError)
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}
	tokens, err := auth.GenerateJWT(user.Username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, tokens)
}

// @Summary     GenerateTokenFromRefreshToken
// @Schemes     http
// @Description Generate a new access and refresh token from an email/password
// @Description  * Refresh token is valid 15 days
// @Description  * Access token is valid 1 hour
// @Tags        Jwt token
// @Produce     json
// @Success     200 {object} auth.JWTTokens
// @Failure     401 {object} object{error=string}
// @Failure     500 {object} object{error=string}
// @Router      /api/token/from-refresh  [post]
// @Security    JWTToken
// @Param       Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func GenerateTokenFromRefreshToken(context *gin.Context) {
	authorizationString := context.GetHeader("Authorization")

	if authorizationString == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "request does not contain an token"})
		context.Abort()
		return
	}

	username, err := auth.ValidateToken(authorizationString, auth.Refresh)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	tokens, err := auth.GenerateJWT(username)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, tokens)
}
