package url

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/w1png/urlshortener/storage"
	"github.com/w1png/urlshortener/utils"
)

func setup() {
  err := utils.ConfigInstance.Init()
  if err != nil {
    panic(err)
  }

  err = storage.InitSelectedStorage()
  if err != nil {
    panic(err)
  }
}

func TestNewUrlService(t *testing.T) {
  s := NewUrlService()
  assert.Equal(t, reflect.TypeOf(s), reflect.TypeOf(&urlService{}))
  
}

func TestCreateGetUrl(t *testing.T) {
  setup()

  s := NewUrlService()
  
  url, err := s.CreateUrl("https://google.com")
  assert.Nil(t, err)
  assert.NotNil(t, url)
  
  url2, err := s.GetUrl(url.Alias)
  assert.Nil(t, err)
  assert.NotNil(t, url2)
  assert.Equal(t, url, url2)
}
