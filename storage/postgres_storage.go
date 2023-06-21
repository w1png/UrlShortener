package storage

import (
	"errors"
	"fmt"
	"os"

	"github.com/w1png/urlshortener/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
  "gorm.io/gorm/logger"
)

type PostgresStorage struct {
	DB *gorm.DB
}

func (ps *PostgresStorage) generateDSN(is_test bool) (string, StorageError) {
  host, ok := os.LookupEnv("POSTGRES_HOST")
	if !ok {
		return "", NewEnvironmentVariableError("POSTGRES_HOST")
	}

  port, ok := os.LookupEnv("POSTGRES_PORT")
	if !ok {
		return "", NewEnvironmentVariableError("POSTGRES_PORT")
	}

	user, ok := os.LookupEnv("POSTGRES_USER")
  if !ok {
		return "", NewEnvironmentVariableError("POSTGRES_USER")
	}

	password, ok := os.LookupEnv("POSTGRES_PASSWORD")
	if !ok {
		return "", NewEnvironmentVariableError("POSTGRES_PASSWORD")
	}

  dbname_var := "POSTGRES_DBNAME"
  if is_test {
    dbname_var = "POSTGRES_TEST_DBNAME"
  }

	dbname, ok := os.LookupEnv(dbname_var)
  if !ok {
		return "", NewEnvironmentVariableError(dbname_var)
	}

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname), nil
}

func (ps *PostgresStorage) connect(dsn string) StorageError {
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
    Logger: logger.Default.LogMode(logger.Silent),
  })
  if err != nil {
    return NewDatabaseConnectionError(err.Error())
  }

  ps.DB = db

  return nil
}

func NewPostgresStorage(is_test bool) (*PostgresStorage, StorageError) {
  dsn, err := (&PostgresStorage{}).generateDSN(is_test)
  if err != nil {
    return nil, err
  }

  ps := &PostgresStorage{}
  err = ps.connect(dsn)
  if err != nil {
    return nil, NewDatabaseConnectionError(err.Error())
  }

  return ps, nil
}

func (ps *PostgresStorage) Save(url *models.Url) StorageError {
  if url == nil {
    return NewUrlIsNilError()
  }

  if ps.DB == nil {
    return NewDatabaseConnectionError("database connection is nil")
  }

  result := ps.DB.Create(url)
  if result.Error != nil {
    return NewDatabaseQueryError(result.Error.Error())
  }

  return nil
}

func (ps *PostgresStorage) GetByAlias(alias string) (*models.Url, StorageError) {
  if alias == "" {
    return nil, NewEmptyAliasError()
  }

  if ps.DB == nil {
    return nil, NewDatabaseConnectionError("database connection is nil")
  }

  url := &models.Url{}
  result := ps.DB.First(url, "alias = ?", alias)
  if result.Error != nil {
    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
      return nil, NewNotFoundError(fmt.Sprintf("url with alias %s", alias))
    }

    return nil, NewDatabaseQueryError(result.Error.Error())
  }

  return url, nil
}

