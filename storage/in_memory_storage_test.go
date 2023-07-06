package storage

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/w1png/urlshortener/models"
	"github.com/w1png/urlshortener/utils"
)

func TestNewInMemoryStorage(t *testing.T) {
	storage := NewInMemoryStorage()
	assert.NotNil(t, storage)
	assert.Equal(t, 0, len(storage.Storage))
	assert.Equal(t, false, storage.Lock)
}

func TestInMemoryStorage_SaveGetUrl(t *testing.T) {
	storage := NewInMemoryStorage()

	url := models.NewUrl("https://google.com")

	err := storage.SaveUrl(url)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(storage.Storage))
	assert.Equal(t, false, storage.Lock)

  url2, err := storage.GetUrlByAlias(url.Alias)
  assert.Nil(t, err)
  assert.Equal(t, url, url2)
  assert.Equal(t, false, storage.Lock)
}

func TestInMemoryStorage_SaveUrl_StorageLockedError(t *testing.T) {
	storage := NewInMemoryStorage()

  url := models.NewUrl("https://google.com")

	storage.Lock = true

	err := storage.SaveUrl(url)
	assert.NotNil(t, err)
	assert.Equal(t, reflect.TypeOf(&utils.StorageLockedError{}), reflect.TypeOf(err))
	assert.Equal(t, 0, len(storage.Storage))
	assert.Equal(t, true, storage.Lock)
}

func TestInMemoryStorage_SaveUrl_UrlIsNilError(t *testing.T) {
	storage := NewInMemoryStorage()

	err := storage.SaveUrl(nil)
	assert.NotNil(t, err)
  assert.Equal(t, reflect.TypeOf(&utils.UrlIsNilError{}), reflect.TypeOf(err))
	assert.Equal(t, 0, len(storage.Storage))
	assert.Equal(t, false, storage.Lock)
}

func TestInMemoryStorage_SaveUrl_EmptyAliasError(t *testing.T) {
	storage := NewInMemoryStorage()

	url := &models.Url{
		Alias: "",
		Url:   "url",
	}

	err := storage.SaveUrl(url)
	assert.NotNil(t, err)
  assert.Equal(t, reflect.TypeOf(&utils.EmptyAliasError{}), reflect.TypeOf(err))
	assert.Equal(t, 0, len(storage.Storage))
	assert.Equal(t, false, storage.Lock)
}

func TestInMemoryStorage_SaveUrl_EmptyUrlError(t *testing.T) {
	storage := NewInMemoryStorage()

	url := &models.Url{
		Alias: "alias",
		Url:   "",
	}

	err := storage.SaveUrl(url)
	assert.NotNil(t, err)
  assert.Equal(t, reflect.TypeOf(&utils.EmptyUrlError{}), reflect.TypeOf(err))
	assert.Equal(t, 0, len(storage.Storage))
	assert.Equal(t, false, storage.Lock)
}

func TestInMemoryStorage_Save_UrlAlreadyExistsError(t *testing.T) {
	storage := NewInMemoryStorage()

	url := &models.Url{
		Alias: "alias",
		Url:   "url",
	}

	storage.Storage[url.Alias] = url.Url

	err := storage.SaveUrl(url)
	assert.NotNil(t, err)
  assert.Equal(t, reflect.TypeOf(&utils.UrlAlreadyExistsError{}), reflect.TypeOf(err))
	assert.Equal(t, 1, len(storage.Storage))
	assert.Equal(t, false, storage.Lock)
}

func TestInMemoryStorage_GetUrlByAlias_NotFoundError(t *testing.T) {
	storage := NewInMemoryStorage()

	result, err := storage.GetUrlByAlias("alias")
	assert.NotNil(t, err)
  assert.Equal(t, reflect.TypeOf(&utils.NotFoundError{}), reflect.TypeOf(err))
	assert.Nil(t, result)
	assert.Equal(t, 0, len(storage.Storage))
	assert.Equal(t, false, storage.Lock)
}

func TestInMemoryStorage_GetUrlByAlias_EmptyAliasError(t *testing.T) {
	storage := NewInMemoryStorage()

	result, err := storage.GetUrlByAlias("")
	assert.NotNil(t, err)
  assert.Equal(t, reflect.TypeOf(&utils.EmptyAliasError{}), reflect.TypeOf(err))
	assert.Nil(t, result)
	assert.Equal(t, 0, len(storage.Storage))
	assert.Equal(t, false, storage.Lock)
}

