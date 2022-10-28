package main

import (
	"fmt"
	"vinpel/golang-sample-api-jwt/common/controllers"
	"vinpel/golang-sample-api-jwt/common/database"
	"vinpel/golang-sample-api-jwt/common/middlewares"
	"vinpel/golang-sample-api-jwt/docs"
	"vinpel/golang-sample-api-jwt/services/ping"
	"vinpel/golang-sample-api-jwt/services/user"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title          golang-sample-api-jwt
// @version        1.0
// @description    Sample API endpoint in golang

// @contact.name  API Support
// @contact.url   https://github.com/vinpel/golang-sample-api-jwt/issues

// @license.name MIT
// @license.url  https://fr.wikipedia.org/wiki/Licence_MIT

// @produce json
// @accept  json

// @host     localhost:8000
// @BasePath /

// @securitydefinitions.apikey JWTToken
// @in                         header
// @name                       Authorization

func main() {
	// Initialize Database
	//get := utils.GetEnvWithKey
	USER := "root"         //get("DB_USER")
	PASS := "root"         //get("DB_PASS")
	HOST := "localhost"    //get("DB_HOST")
	PORT := "3306"         //get("DB_PORT")
	DBNAME := "myDatabase" // get("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT)
	database.Connect(dsn, DBNAME)
	database.Migrate()

	docs.SwaggerInfo.BasePath = "/"

	// Initialize Router
	router := initRouter()

	router.Run("localhost:8000")
}

func initRouter() *gin.Engine {
	router := gin.Default()

	if gin.Mode() != gin.ReleaseMode {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	api := router.Group("/api")
	{
		token := api.Group("/token")
		{
			token.POST("/from-user", controllers.GenerateTokenFromUser)
			token.POST("/from-refresh", controllers.GenerateTokenFromRefreshToken)
		}

		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", ping.Ping)
		}
	}
	user.RegisterRoutes(api)

	return router
}
