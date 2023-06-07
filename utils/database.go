package utils

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func CreateDSN(is_test bool) (string, error) {
	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		return "", fmt.Errorf("POSTGRES_HOST is not set")
	}

	port := os.Getenv("POSTGRES_PORT")
	if port == "" {
		return "", fmt.Errorf("POSTGRES_PORT is not set")
	}

  var db_name_var string
  
  if is_test {
    db_name_var = "POSTGRES_DBNAME_TEST"
  } else {
    db_name_var = "POSTGRES_DBNAME"
  }

	db_name := os.Getenv(db_name_var)
	if db_name == "" {
		return "", fmt.Errorf("%s is not set", db_name_var)
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

func connectDB(is_test bool) (*gorm.DB, error) {
	dsn, err := CreateDSN(is_test)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
    Logger: logger.Default.LogMode(logger.Silent),
  })

	if err != nil {
		return nil, err
	}

	return db, nil
}

func InitDB(is_test bool) error {
	db, err := connectDB(is_test)
	if err != nil {
		return err
	}

	DB = db
	return nil
}
