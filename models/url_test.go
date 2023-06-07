package models

import (
  "testing"
  "github.com/stretchr/testify/assert"

  "github.com/w1png/ozontest/utils"
)

func TestUrl(t *testing.T) {
  utils.InitDB(true)
  utils.DB.AutoMigrate(&Url{})

  url, err := NewUrl("http://www.google.com")
  assert.Nil(t, err)
  assert.NotNil(t, url.Url)
  assert.NotNil(t, url.Alias)

  err = url.Save()
  assert.Nil(t, err)
  assert.NotNil(t, url.Url)
  assert.NotNil(t, url.Alias)

  newUrl, err := GetUrlByAlias(url.Alias)
  assert.Nil(t, err)
  assert.Equal(t, url.Url, newUrl.Url)
  assert.Equal(t, url.Alias, newUrl.Alias)

  newUrl, err = GetUrlByAlias("thisisnotanalias")
  assert.Error(t, err)
  assert.Nil(t, newUrl)
}

