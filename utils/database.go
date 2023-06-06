package utils

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func CreateDSN() (string, error) {
	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		return "", fmt.Errorf("POSTGRES_HOST is not set")
	}

	port := os.Getenv("POSTGRES_PORT")
	if port == "" {
		return "", fmt.Errorf("POSTGRES_PORT is not set")
	}

	db_name := os.Getenv("POSTGRES_DBNAME")
	if db_name == "" {
		return "", fmt.Errorf("POSTGRES_DBNAME is not set")
	}

	user := os.Getenv("POSTGRES_USER")
	if user == "" {
		return "", fmt.Errorf("POSTGRES_USER is not set")
	}

	password := os.Getenv("POSTGRES_PASSWORD")
	if password == "" {
		return "", fmt.Errorf("POSTGRES_PASSWORD is not set")
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, db_name)
	return dsn, nil
}

func ConnectDB() (*gorm.DB, error) {
	dsn, err := CreateDSN()
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func InitDB() error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}

	DB = db
	return nil
}
