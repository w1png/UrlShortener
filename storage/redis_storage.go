package storage

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/w1png/urlshortener/models"
	"github.com/w1png/urlshortener/utils"
)

type RedisStorage struct {
	client *redis.Client
}

func NewRedisStorage(is_test bool) (*RedisStorage, error) {
	db := utils.ConfigInstance.RedisDatabase
	if is_test {
		db = utils.ConfigInstance.RedisTestDatabase
	}

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", utils.ConfigInstance.RedisHost, utils.ConfigInstance.RedisPort),
		Password: utils.ConfigInstance.RedisPassword,
		DB:       db,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, utils.NewDatabaseConnectionError(err.Error())
	}

	return &RedisStorage{client: client}, nil
}

func (s *RedisStorage) SaveUrl(url *models.Url) error {
	ctx := context.Background()
	err := s.client.Set(ctx, url.Alias, url.Url, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (s *RedisStorage) GetUrlByAlias(alias string) (*models.Url, error) {
	ctx := context.Background()
	url, err := s.client.Get(ctx, alias).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, utils.NewNotFoundError(fmt.Sprintf("Url with alias %s", alias))
		}
		return nil, err

	}
	return &models.Url{Alias: alias, Url: url}, nil
}
