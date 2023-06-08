package models

import (
  "testing"
  "github.com/stretchr/testify/assert"

)

func TestNewUrl(t *testing.T) {
  url := NewUrl("https://google.com")
  assert.NotNil(t, url)
  assert.Equal(t, "https://google.com", url.Url)
}

