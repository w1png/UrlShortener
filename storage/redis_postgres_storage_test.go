package storage

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/w1png/urlshortener/models"
	"github.com/w1png/urlshortener/utils"
)

func setupPostgresRedisStorage() {
  err := utils.ConfigInstance.Init()
  if err != nil {
    panic(err)
  }
}

func TestNewPostgresRedisStorage(t *testing.T) {
  setupPostgresRedisStorage()

  s, err := NewRedisPostgresStorage(true)
  assert.Nil(t, err)
  assert.Equal(t, reflect.TypeOf(s), reflect.TypeOf(&RedisPostgresStorage{}))
}

func TestNewPostgresRedisStorage_DatabaseConnectionErrorPostgres(t *testing.T) {
  setupPostgresRedisStorage()

  utils.ConfigInstance.PostgresHost = "wrong_host"
  _, err := NewRedisPostgresStorage(true)
  assert.NotNil(t, err)
}

func TestNewPostgresRedisStorage_DatabaseConnectionErrorRedis(t *testing.T) {
  setupPostgresRedisStorage()

  utils.ConfigInstance.RedisHost = "wrong_host"
  _, err := NewRedisPostgresStorage(true)
  assert.NotNil(t, err)
}

func TestRedisPostgresStorage_SaveUrl(t *testing.T) {
  setupPostgresRedisStorage()

  s, err := NewRedisPostgresStorage(true)
  assert.Nil(t, err)

  url := models.NewUrl("https://google.com")

  err = s.SaveUrl(url)
  assert.Nil(t, err)
}

func TestRedisPostgresStorage_GetUrlByAlias(t *testing.T) {
  setupPostgresRedisStorage()

  s, err := NewRedisPostgresStorage(true)
  assert.Nil(t, err)

  url := models.NewUrl("https://google.com")

  err = s.SaveUrl(url)
  assert.Nil(t, err)

  url2, err := s.GetUrlByAlias(url.Alias)
  assert.Nil(t, err)
  assert.Equal(t, url.Alias, url2.Alias)
  assert.Equal(t, url.Url, url2.Url)
}

func TestRedisPostgresStorage_GetUrlByAlias_NotFound(t *testing.T) {
  setupPostgresRedisStorage()

  s, err := NewRedisPostgresStorage(true)
  assert.Nil(t, err)

  _, err = s.GetUrlByAlias("not_found")
  assert.NotNil(t, err)

  assert.Equal(t, reflect.TypeOf(err), reflect.TypeOf(&utils.NotFoundError{}))
}

