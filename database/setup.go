package database

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	dsn := os.Getenv("DATABASE_DSN")
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	DB = db
}
