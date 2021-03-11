package database

import (
	"../models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "postgres://postgres:1234@localhost:5432/gotutorial"
    connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
    	panic("failed to connect database")
    }

	DB = connection

	connection.AutoMigrate(&models.User{})
}