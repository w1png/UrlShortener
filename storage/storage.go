package storage

import (
	"os"

	"github.com/w1png/urlshortener/models"
	"github.com/w1png/urlshortener/utils"
)

var SelectedStorage Storage

func InitSelectedStorage() error {
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
    return utils.NewEnvironmentVariableError("STORAGE_TYPE")
	}

	return nil
}

type Storage interface {
	SaveUrl(url *models.Url) error
	GetUrlByAlias(alias string) (*models.Url, error)
}
