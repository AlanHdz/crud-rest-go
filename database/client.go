package database

import (
	"alanhedz/golang-crud/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var err error

func Connect(connection string) {
	Instance, err = gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database...")
}

func Migrate() {
	Instance.AutoMigrate(&models.User{})
	log.Println("Database migration completed...")
}
