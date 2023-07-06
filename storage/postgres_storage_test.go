package storage

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/w1png/urlshortener/models"
	"github.com/w1png/urlshortener/utils"
)

func setup() {
  err := utils.ConfigInstance.Init()
  if err != nil {
    panic(err)
  }
}

func TestNewPostgresStorage(t *testing.T) {
  setup()

  storage, err := NewPostgresStorage(true)
  assert.NotNil(t, storage)
  assert.Nil(t, err)
}

func TestPostgresStorage_NewPostgresStorage_DatabaseConnectionError(t *testing.T) {
  setup()

  utils.ConfigInstance.PostgresHost = "invalid host"
  storage, err := NewPostgresStorage(true)
  assert.Nil(t, storage)
  assert.NotNil(t, err)
  assert.Equal(t, reflect.TypeOf(err), reflect.TypeOf(&utils.DatabaseConnectionError{}))
}

func TestPostgresStorage_SaveGetUrl(t *testing.T) {
  setup()

  storage, err := NewPostgresStorage(true)
  assert.NotNil(t, storage)
  assert.Nil(t, err)

  url := models.NewUrl("https://google.com")
  err = storage.SaveUrl(url)
  assert.Nil(t, err)

  url2, err := storage.GetUrlByAlias(url.Alias)
  assert.Nil(t, err)
  assert.Equal(t, url, url2)
}

func TestPostgresStorage_SaveUrl_UrlIsNilError(t *testing.T) {
  setup()

  storage, err := NewPostgresStorage(true)
  assert.NotNil(t, storage)
  assert.Nil(t, err)

  err = storage.SaveUrl(nil)
  assert.NotNil(t, err)
  assert.Equal(t, reflect.TypeOf(err), reflect.TypeOf(&utils.UrlIsNilError{}))
}

func TestPostgresStorage_SaveUrl_DatabaseConnectionError(t *testing.T) {
  setup()

  storage, err := NewPostgresStorage(true)
  assert.NotNil(t, storage)
  assert.Nil(t, err)

  storage.DB = nil
  url := models.NewUrl("https://google.com")
  err = storage.SaveUrl(url)
  assert.NotNil(t, err)
  assert.Equal(t, reflect.TypeOf(err), reflect.TypeOf(&utils.DatabaseConnectionError{}))
}

func TestPostgresStorage_GetUrlByAlias_EmptyAliasError(t *testing.T) {
  setup()

  storage, err := NewPostgresStorage(true)
  assert.NotNil(t, storage)
  assert.Nil(t, err)

  _, err = storage.GetUrlByAlias("")
  assert.NotNil(t, err)
  assert.Equal(t, reflect.TypeOf(err), reflect.TypeOf(&utils.EmptyAliasError{}))
}

func TestPostgresStorage_GetUrlByAlias_DatabaseConnectionError(t *testing.T) {
  setup()

  storage, err := NewPostgresStorage(true)
  assert.NotNil(t, storage)
  assert.Nil(t, err)

  storage.DB = nil
  _, err = storage.GetUrlByAlias("alias")
  assert.NotNil(t, err)
  assert.Equal(t, reflect.TypeOf(err), reflect.TypeOf(&utils.DatabaseConnectionError{}))
}

func TestPostgresStorage_GetUrlByAlias_NotFoundError(t *testing.T) {
  setup()

  storage, err := NewPostgresStorage(true)
  assert.NotNil(t, storage)
  assert.Nil(t, err)

  _, err = storage.GetUrlByAlias("alias")
  assert.NotNil(t, err)
  assert.Equal(t, reflect.TypeOf(err), reflect.TypeOf(&utils.NotFoundError{}))
}

