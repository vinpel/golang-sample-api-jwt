package database

import (
	"log"
	"vinpel/golang-sample-api-jwt/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string, dbname string) {
	if Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{}); dbError != nil {
		log.Fatal(dbError)
	}
	if result := Instance.Exec("CREATE DATABASE IF NOT EXISTS " + dbname + ";"); result.Error != nil {
		log.Fatal(result.Error)

	}
	if result := Instance.Exec("USE " + dbname); result.Error != nil {
		log.Fatal(result.Error)
	}
	log.Println("Connected to Database!")
}
func Migrate() {
	Instance.AutoMigrate(&models.User{})
	log.Println("Database Migration Completed!")
}
