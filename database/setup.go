// Package database provides functions to connect to a MySQL database using GORM.
package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB represents the global database connection.
var DB *gorm.DB

// ConnectDatabase establishes a connection to the MySQL database using the provided DSN.
// It sets the global DB variable to the established connection.
// If an error occurs during the connection process, it is returned.
func ConnectDatabase() error {
	dsn := os.Getenv("DATABASE_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}
	DB = db
	return nil
}
