package storage

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/w1png/urlshortener/models"
)

func TestNewPostgresStorage(t *testing.T) {
  storage, err := NewPostgresStorage(true)
  assert.NotNil(t, storage)
  assert.Nil(t, err)
}

func TestPostgresStorage_NewPostgresStorage_EnviromentVariableError(t *testing.T) {
  host := os.Getenv("POSTGRES_HOST")
  port := os.Getenv("POSTGRES_PORT")
  user := os.Getenv("POSTGRES_USER")
  password := os.Getenv("POSTGRES_PASSWORD")
  database := os.Getenv("POSTGRES_TEST_DBNAME")
  defer func(host, port, user, password, database string) {
    os.Setenv("POSTGRES_HOST", host)
    os.Setenv("POSTGRES_PORT", port)
    os.Setenv("POSTGRES_USER", user)
    os.Setenv("POSTGRES_PASSWORD", password)
    os.Setenv("POSTGRES_TEST_DBNAME", database)
  }(host, port, user, password, database)

  testEnvVariableError := func(env string) {
    os.Setenv(env, "")
    storage, err := NewPostgresStorage(true)
    assert.Nil(t, storage)
    assert.NotNil(t, err)
    assert.Equal(t, reflect.TypeOf(err), reflect.TypeOf(&EnvironmentVariableError{}))
  }

  testEnvVariableError("POSTGRES_HOST")
  testEnvVariableError("POSTGRES_PORT")
  testEnvVariableError("POSTGRES_USER")
  testEnvVariableError("POSTGRES_PASSWORD")
  testEnvVariableError("POSTGRES_TEST_DBNAME")
}

func TestPostgresStorage_NewPostgresStorage_DatabaseConnectionError(t *testing.T) {
  host := os.Getenv("POSTGRES_HOST")
  defer os.Setenv("POSTGRES_HOST", host)

  os.Setenv("POSTGRES_HOST", "wrong_host")
  storage, err := NewPostgresStorage(true)
  assert.Nil(t, storage)
  assert.NotNil(t, err)
  assert.Equal(t, reflect.TypeOf(err), reflect.TypeOf(&DatabaseConnectionError{}))
}

func TestPostgresStorage_SaveAndGet(t *testing.T) {
  storage, err := NewPostgresStorage(true)
  assert.NotNil(t, storage)
  assert.Nil(t, err)

  url := models.NewUrl("https://google.com")
  err = storage.Save(url)
  assert.Nil(t, err)

  url2, err := storage.GetByAlias(url.Alias)
  assert.Nil(t, err)
  assert.Equal(t, url, url2)
}

func TestPostgresStorage_Save_UrlIsNilError(t *testing.T) {
  storage, err := NewPostgresStorage(true)
  assert.NotNil(t, storage)
  assert.Nil(t, err)

  err = storage.Save(nil)
  assert.NotNil(t, err)
  assert.Equal(t, reflect.TypeOf(err), reflect.TypeOf(&UrlIsNilError{}))
}

func TestPostgresStorage_Save_DatabaseConnectionError(t *testing.T) {
  storage, err := NewPostgresStorage(true)
  assert.NotNil(t, storage)
  assert.Nil(t, err)

  storage.DB = nil
  url := models.NewUrl("https://google.com")
  err = storage.Save(url)
  assert.NotNil(t, err)
  assert.Equal(t, reflect.TypeOf(err), reflect.TypeOf(&DatabaseConnectionError{}))
}

func TestPostgresStorage_GetByAlias_EmptyAliasError(t *testing.T) {
  storage, err := NewPostgresStorage(true)
  assert.NotNil(t, storage)
  assert.Nil(t, err)

  _, err = storage.GetByAlias("")
  assert.NotNil(t, err)
  assert.Equal(t, reflect.TypeOf(err), reflect.TypeOf(&EmptyAliasError{}))
}

func TestPostgresStorage_GetByAlias_DatabaseConnectionError(t *testing.T) {
  storage, err := NewPostgresStorage(true)
  assert.NotNil(t, storage)
  assert.Nil(t, err)

  storage.DB = nil
  _, err = storage.GetByAlias("alias")
  assert.NotNil(t, err)
  assert.Equal(t, reflect.TypeOf(err), reflect.TypeOf(&DatabaseConnectionError{}))
}

func TestPostgresStorage_GetByAlias_NotFoundError(t *testing.T) {
  storage, err := NewPostgresStorage(true)
  assert.NotNil(t, storage)
  assert.Nil(t, err)

  _, err = storage.GetByAlias("alias")
  assert.NotNil(t, err)
  assert.Equal(t, reflect.TypeOf(err), reflect.TypeOf(&NotFoundError{}))
}

