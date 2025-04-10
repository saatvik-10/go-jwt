package database

import (
	"jwt/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:Messi@1987@/go_jwt"), &gorm.Config{})

	if err != nil {
		panic("Database connection failed")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
