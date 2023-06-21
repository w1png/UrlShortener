package url

import (
	"github.com/w1png/urlshortener/models"
	"github.com/w1png/urlshortener/storage"
)

type urlService struct {}

func (s *urlService) CreateUrl(urlBase string) (Url, storage.StorageError) {
  url := models.NewUrl(urlBase)
  err := storage.SelectedStorage.Save(url)
  
  if err != nil {
    return Url{}, err
  }
  return Url{url.Url, url.Alias}, nil
}

func (s *urlService) GetUrl(alias string) (Url, storage.StorageError) {
  url, err := storage.SelectedStorage.GetByAlias(alias)
  if err != nil {
    return Url{}, err
  }
  return Url{url.Url, url.Alias}, nil
}

func NewUrlService() Service {
  return &urlService{}
}
