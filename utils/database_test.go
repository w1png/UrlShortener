package utils

import (
  "testing"
  "github.com/stretchr/testify/assert"


)

func TestInitTestDB(t *testing.T) {
  err := InitDB(true)
  assert.NoError(t, err)
}

func TestInitDB(t *testing.T) {
  err := InitDB(false)
  assert.NoError(t, err)
}
