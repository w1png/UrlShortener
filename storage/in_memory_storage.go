package storage

import (
	"fmt"

	"github.com/w1png/urlshortener/models"
	"github.com/w1png/urlshortener/utils"
)

type InMemoryStorage struct{
  Lock bool
  Storage map[string]string
}

func NewInMemoryStorage() *InMemoryStorage {
  return &InMemoryStorage{
    Lock: false,
    Storage: make(map[string]string),
  }
}

func (s *InMemoryStorage) SaveUrl(url *models.Url) error {
  if s.Lock {
    return utils.NewStorageLockedError()
  }

  s.Lock = true
  defer func() {
    s.Lock = false
  }()

  if url == nil {
    return utils.NewUrlIsNilError()
  }
  if url.Alias == "" {
    return utils.NewEmptyAliasError()
  }
  if url.Url == "" {
    return utils.NewEmptyUrlError()
  }

  if _, ok := s.Storage[url.Alias]; ok {
    return utils.NewUrlAlreadyExistsError(fmt.Sprintf("url with alias %s", url.Alias))
  }

  s.Storage[url.Alias] = url.Url

  return nil
}

func (s *InMemoryStorage) GetUrlByAlias(alias string) (*models.Url, error) {
  if alias == "" {
    return nil, utils.NewEmptyAliasError()
  }

  if url, ok := s.Storage[alias]; ok {
    return &models.Url{
      Alias: alias,
      Url:   url,
    }, nil
  } 

  return nil, utils.NewNotFoundError(fmt.Sprintf("url with alias %s not found", alias))
}

