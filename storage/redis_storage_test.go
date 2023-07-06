package storage

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/w1png/urlshortener/models"
	"github.com/w1png/urlshortener/utils"
)

func setupRedis() {
  err := utils.ConfigInstance.Init()
  if err != nil {
    panic(err)
  }
}

func TestNewRedisStorage(t *testing.T) {
  setupRedis()

  storage, err := NewRedisStorage(true)
  assert.Nil(t, err)
  assert.NotNil(t, storage)
}

func TestNewRedisStorage_DatabaseConnectionError(t *testing.T) {
  setupRedis()
  utils.ConfigInstance.RedisHost = "wrong_host"

  storage, err := NewRedisStorage(false)
  assert.NotNil(t, err)
  assert.Nil(t, storage)
}

func TestRedisStorage_SaveGetUrl(t *testing.T) {
  setupRedis()
  storage, err := NewRedisStorage(true)
  assert.Nil(t, err)
  assert.NotNil(t, storage)

  url := models.NewUrl("https://google.com")
  err = storage.SaveUrl(url)
  assert.Nil(t, err)

  url2, err := storage.GetUrlByAlias(url.Alias)
  assert.Nil(t, err)
  assert.Equal(t, url, url2)
}

func TestRedisStorage_NotFoundError(t *testing.T) {
  setupRedis()
  storage, err := NewRedisStorage(true)
  assert.Nil(t, err)
  assert.NotNil(t, storage)

  _, err = storage.GetUrlByAlias("invalid alias")
  assert.NotNil(t, err)
  assert.Equal(t, reflect.TypeOf(err), reflect.TypeOf(&utils.NotFoundError{}))
}

