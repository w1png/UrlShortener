package storage

import (
	"fmt"

	"github.com/w1png/ozontask/models"
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

func (s *InMemoryStorage) Save(url *models.Url) StorageError {
  if s.Lock {
    return NewStorageLockedError("storage is locked")
  }

  s.Lock = true
  defer func() {
    s.Lock = false
  }()

  if url == nil {
    return NewUrlIsNilError("url is nil")
  }
  if url.Alias == "" {
    return NewEmptyAliasError("alias is empty")
  }
  if url.Url == "" {
    return NewEmptyUrlError("url is empty")
  }

  if _, ok := s.Storage[url.Alias]; ok {
    return NewUrlAlreadyExistsError(fmt.Sprintf("url with alias %s already exists", url.Alias))
  }

  s.Storage[url.Alias] = url.Url

  return nil
}

func (s *InMemoryStorage) GetByAlias(alias string) (*models.Url, StorageError) {
  if alias == "" {
    return nil, NewEmptyAliasError("alias is empty")
  }

  if url, ok := s.Storage[alias]; ok {
    return &models.Url{
      Alias: alias,
      Url:   url,
    }, nil
  } 

  return nil, NewNotFoundError(fmt.Sprintf("url with alias %s not found", alias))
}

