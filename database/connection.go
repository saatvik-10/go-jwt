package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	_, err := gorm.Open(mysql.Open("root:Messi@1987@/go_jwt"), &gorm.Config{})

	if err != nil {
		panic("Database connection failed")
	}
}
