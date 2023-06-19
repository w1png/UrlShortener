package url

import (
	"fmt"

	"github.com/w1png/urlshortener/models"
	"github.com/w1png/urlshortener/storage"
)

type urlService struct {}

func (s *urlService) CreateUrl(urlBase string) (Url, error) {
  fmt.Println("createUrl with url: ", urlBase)
  url := models.NewUrl(urlBase)
  storage.SelectedStorage.Save(url)
  fmt.Printf("url: %+v\n", url)
  return Url{url.Url, url.Alias}, nil
}

func (s *urlService) GetUrl(alias string) (Url, error) {
  fmt.Println("getUrl with alias: ", alias)
  url, err := storage.SelectedStorage.GetByAlias(alias)
  if err != nil {
    return Url{}, err
  }
  return Url{url.Url, url.Alias}, nil
}

func NewUrlService() Service {
  return &urlService{}
}
