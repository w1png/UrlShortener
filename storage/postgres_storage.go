package storage

import (
	"errors"
	"fmt"

	"github.com/w1png/urlshortener/models"
	"github.com/w1png/urlshortener/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresStorage struct {
	DB *gorm.DB
}

func (ps *PostgresStorage) generateDSN(is_test bool) (string, error) {
	dbname := utils.ConfigInstance.PostgresDatabase
  if is_test {
    dbname = utils.ConfigInstance.PostgresTestDatabase
  }

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", utils.ConfigInstance.PostgresHost, utils.ConfigInstance.PostgresPort, utils.ConfigInstance.PostgresUser, utils.ConfigInstance.PostgresPassword, dbname), nil
}

func (ps *PostgresStorage) connect(dsn string) error {
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
    Logger: logger.Default.LogMode(logger.Silent),
  })
  if err != nil {
    return utils.NewDatabaseConnectionError(err.Error())
  }

  ps.DB = db

  return nil
}

func NewPostgresStorage(is_test bool) (*PostgresStorage, error) {
  dsn, err := (&PostgresStorage{}).generateDSN(is_test)
  if err != nil {
    return nil, err
  }

  ps := &PostgresStorage{}
  err = ps.connect(dsn)
  if err != nil {
    return nil, utils.NewDatabaseConnectionError(err.Error())
  }

  return ps, nil
}

func (ps *PostgresStorage) SaveUrl(url *models.Url) error {
  if url == nil {
    return utils.NewUrlIsNilError()
  }

  if ps.DB == nil {
    return utils.NewDatabaseConnectionError("database connection is nil")
  }

  result := ps.DB.Create(url)
  if result.Error != nil {
    return utils.NewDatabaseQueryError(result.Error.Error())
  }

  return nil
}

func (ps *PostgresStorage) GetUrlByAlias(alias string) (*models.Url, error) {
  if alias == "" {
    return nil, utils.NewEmptyAliasError()
  }

  if ps.DB == nil {
    return nil, utils.NewDatabaseConnectionError("database connection is nil")
  }

  url := &models.Url{}
  result := ps.DB.First(url, "alias = ?", alias)
  if result.Error != nil {
    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
      return nil, utils.NewNotFoundError(fmt.Sprintf("url with alias %s", alias))
    }

    return nil, utils.NewDatabaseQueryError(result.Error.Error())
  }

  return url, nil
}

