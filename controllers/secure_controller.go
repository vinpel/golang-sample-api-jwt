package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Ping
// @Schemes http
// @Description respond pong to a ping with a JWT Header
// @Tags     Jwt token
// @Produce  json
// @Success  200 {object} object{message=string}
// @Failure  401 {object} object{error=string}
// @Failure  500 {object} object{error=string}
// @Router   /api/secured/ping  [get]
// @Security JWTToken
// @Param    Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func Ping(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "pong"})
}
