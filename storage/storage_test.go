package storage

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/w1png/urlshortener/utils"
)

func TestInitSelectedStorage_InMemory(t *testing.T) {
  storage_type := os.Getenv("STORAGE_TYPE")
  defer os.Setenv("STORAGE_TYPE", storage_type)

  os.Setenv("STORAGE_TYPE", "in_memory")
  err := InitSelectedStorage()
  assert.Nil(t, err)
  assert.IsType(t, &InMemoryStorage{}, SelectedStorage)
  assert.Equal(t, reflect.TypeOf(&InMemoryStorage{}), reflect.TypeOf(SelectedStorage))
}

func TestInitSelectedStorage_Postgres(t *testing.T) {
  storage_type := os.Getenv("STORAGE_TYPE")
  defer os.Setenv("STORAGE_TYPE", storage_type)

  os.Setenv("STORAGE_TYPE", "postgres")
  err := InitSelectedStorage()
  assert.Nil(t, err)
  assert.IsType(t, &PostgresStorage{}, SelectedStorage)
  assert.Equal(t, reflect.TypeOf(&PostgresStorage{}), reflect.TypeOf(SelectedStorage))
}

func TestInitSelectedStorage_EnvironmentVariableError(t *testing.T) {
  storage_type := os.Getenv("STORAGE_TYPE")
  defer os.Setenv("STORAGE_TYPE", storage_type)

  os.Unsetenv("STORAGE_TYPE")
  err := InitSelectedStorage()
  assert.NotNil(t, err)
  assert.Equal(t, "Environment variable error: STORAGE_TYPE", err.Error())
  assert.Equal(t, reflect.TypeOf(&utils.EnvironmentVariableError{}), reflect.TypeOf(err))
}
