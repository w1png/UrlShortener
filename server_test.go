package main

import (
  "testing"

  "github.com/stretchr/testify/assert"

)

func TestNewApiServer(t *testing.T) {
  server := NewApiServer(":8080")
  assert.NotNil(t, server)

  assert.Equal(t, ":8080", server.listenAddr)
}
