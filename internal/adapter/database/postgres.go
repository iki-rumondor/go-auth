package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB() (*gorm.DB, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	sslMode := "disable"

	strConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname%s sslmode=%s", dbHost, dbPort, dbUser, dbPassword, dbName, sslMode)

	gormDB, err := gorm.Open(postgres.Open(strConn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return gormDB, nil
}
