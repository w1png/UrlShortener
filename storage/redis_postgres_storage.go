package storage

import (
	"github.com/w1png/urlshortener/models"
)

type RedisPostgresStorage struct {
  redis_storage *RedisStorage
  postgres_storage *PostgresStorage
}

func NewRedisPostgresStorage(test bool) (*RedisPostgresStorage, error) {
  redis_storage, err := NewRedisStorage(test)
  if err != nil {
    return nil, err
  }

  postgres_storage, err := NewPostgresStorage(test)
  if err != nil {
    return nil, err
  }

  return &RedisPostgresStorage{
    redis_storage: redis_storage,
    postgres_storage: postgres_storage,
  }, nil
}

func (s *RedisPostgresStorage) SaveUrl(url *models.Url) error {
  err := s.postgres_storage.SaveUrl(url)
  if err != nil {
    return err
  }

  err = s.redis_storage.SaveUrl(url)
  if err != nil {
    return err
  }
  
  return nil
}

func (s *RedisPostgresStorage) GetUrlByAlias(alias string) (*models.Url, error) {
  url, err := s.redis_storage.GetUrlByAlias(alias)
  if err == nil {
    return url, nil
  }

  url, err = s.postgres_storage.GetUrlByAlias(alias)
  if err != nil {
    return nil, err
  }

  err = s.redis_storage.SaveUrl(url)
  if err != nil {
    return nil, err
  }

  return url, nil
}

