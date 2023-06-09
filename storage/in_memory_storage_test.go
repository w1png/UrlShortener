package storage

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/w1png/urlshortener/models"
)

func TestNewInMemoryStorage(t *testing.T) {
	storage := NewInMemoryStorage()
	assert.NotNil(t, storage)
	assert.Equal(t, 0, len(storage.Storage))
	assert.Equal(t, false, storage.Lock)
}

func TestInMemoryStorage_Save(t *testing.T) {
	storage := NewInMemoryStorage()

	url := models.NewUrl("https://google.com")

	err := storage.Save(url)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(storage.Storage))
	assert.Equal(t, false, storage.Lock)
}

func TestInMemoryStorage_Save_StorageLockedError(t *testing.T) {
	storage := NewInMemoryStorage()

  url := models.NewUrl("https://google.com")

	storage.Lock = true

	err := storage.Save(url)
	assert.NotNil(t, err)
	assert.Equal(t, reflect.TypeOf(&StorageLockedError{}), reflect.TypeOf(err))
	assert.Equal(t, 0, len(storage.Storage))
	assert.Equal(t, true, storage.Lock)
}

func TestInMemoryStorage_Save_UrlIsNilError(t *testing.T) {
	storage := NewInMemoryStorage()

	err := storage.Save(nil)
	assert.NotNil(t, err)
  assert.Equal(t, reflect.TypeOf(&UrlIsNilError{}), reflect.TypeOf(err))
	assert.Equal(t, 0, len(storage.Storage))
	assert.Equal(t, false, storage.Lock)
}

func TestInMemoryStorage_Save_EmptyAliasError(t *testing.T) {
	storage := NewInMemoryStorage()

	url := &models.Url{
		Alias: "",
		Url:   "url",
	}

	err := storage.Save(url)
	assert.NotNil(t, err)
  assert.Equal(t, reflect.TypeOf(&EmptyAliasError{}), reflect.TypeOf(err))
	assert.Equal(t, 0, len(storage.Storage))
	assert.Equal(t, false, storage.Lock)
}

func TestInMemoryStorage_Save_EmptyUrlError(t *testing.T) {
	storage := NewInMemoryStorage()

	url := &models.Url{
		Alias: "alias",
		Url:   "",
	}

	err := storage.Save(url)
	assert.NotNil(t, err)
  assert.Equal(t, reflect.TypeOf(&EmptyUrlError{}), reflect.TypeOf(err))
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

	err := storage.Save(url)
	assert.NotNil(t, err)
  assert.Equal(t, reflect.TypeOf(&UrlAlreadyExistsError{}), reflect.TypeOf(err))
	assert.Equal(t, 1, len(storage.Storage))
	assert.Equal(t, false, storage.Lock)
}

func TestInMemoryStorage_GetByAlias(t *testing.T) {
	storage := NewInMemoryStorage()

	url := &models.Url{
		Alias: "alias",
		Url:   "url",
	}

	storage.Storage[url.Alias] = url.Url

	result, err := storage.GetByAlias(url.Alias)
	assert.Nil(t, err)
	assert.Equal(t, url, result)
	assert.Equal(t, false, storage.Lock)
}

func TestInMemoryStorage_GetByAlias_NotFoundError(t *testing.T) {
	storage := NewInMemoryStorage()

	result, err := storage.GetByAlias("alias")
	assert.NotNil(t, err)
  assert.Equal(t, reflect.TypeOf(&NotFoundError{}), reflect.TypeOf(err))
	assert.Nil(t, result)
	assert.Equal(t, 0, len(storage.Storage))
	assert.Equal(t, false, storage.Lock)
}

func TestInMemoryStorage_GetByAlias_EmptyAliasError(t *testing.T) {
	storage := NewInMemoryStorage()

	result, err := storage.GetByAlias("")
	assert.NotNil(t, err)
  assert.Equal(t, reflect.TypeOf(&EmptyAliasError{}), reflect.TypeOf(err))
	assert.Nil(t, result)
	assert.Equal(t, 0, len(storage.Storage))
	assert.Equal(t, false, storage.Lock)
}

