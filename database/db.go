package database

import (
	"log"

	"github.com/luciormoraes/gin-rest-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {
	dsn := "host=172.20.0.2 user=root password=root dbname=gin_api port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Panic("Connection to DB failing")
	}
	DB.AutoMigrate(&models.Student{})
}
