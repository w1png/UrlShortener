package url

import (
	"github.com/w1png/urlshortener/models"
	"github.com/w1png/urlshortener/storage"
)

type urlService struct {}

func (s *urlService) CreateUrl(urlBase string) (Url, error) {
  url := models.NewUrl(urlBase)
  storage.SelectedStorage.Save(url)
  return Url{url.Url, url.Alias}, nil
}

func (s *urlService) GetUrl(alias string) (Url, error) {
  url, err := storage.SelectedStorage.GetByAlias(alias)
  if err != nil {
    return Url{}, err
  }
  return Url{url.Url, url.Alias}, nil
}

func NewUrlService() Service {
  return &urlService{}
}
