package storage

import (
	"os"

	"github.com/w1png/urlshortener/models"
)

var SelectedStorage UrlStorage

func InitSelectedStorage() StorageError {
	switch os.Getenv("STORAGE_TYPE") {
	case "in_memory":
		SelectedStorage = NewInMemoryStorage()
	case "postgres":
    var err error
		SelectedStorage, err = NewPostgresStorage(false)
		if err != nil {
			return err
		}
  default:
    return NewEnvironmentVariableError("STORAGE_TYPE")
	}

	return nil
}

type UrlStorage interface {
	Save(url *models.Url) StorageError
	GetByAlias(alias string) (*models.Url, StorageError)
}
