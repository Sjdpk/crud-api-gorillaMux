package database

import (
	"golang-crud/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() {
	dsn := "host=localhost user=iamdpk password=iamdpk dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Kathmandu"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	Database = database
	database.AutoMigrate(&models.BooksModel{})
}
