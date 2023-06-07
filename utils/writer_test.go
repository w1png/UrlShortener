package utils

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestObject struct {
  Url string `json:"url"`
  Alias string `json:"alias"`
}

func TestWriteResponse(t *testing.T) {
  w := httptest.NewRecorder()
  WriteResponse(w, http.StatusOK, "test")
  assert.Equal(t, http.StatusOK, w.Code)

  testObj := TestObject{Url: "http://www.google.com", Alias: "google"}
  w = httptest.NewRecorder()
  WriteResponse(w, http.StatusCreated, testObj)
  assert.Equal(t, http.StatusCreated, w.Code)

  w = httptest.NewRecorder()
  WriteResponse(w, http.StatusNotFound, nil)
  assert.Equal(t, http.StatusNotFound, w.Code)

  w = httptest.NewRecorder()
  WriteResponse(w, http.StatusInternalServerError, nil)
  assert.Equal(t, http.StatusInternalServerError, w.Code)

  w = httptest.NewRecorder()
  WriteResponse(w, http.StatusBadRequest, nil)
  assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestWriteError(t *testing.T) {
  w := httptest.NewRecorder()
  WriteError(w, http.StatusNotFound, fmt.Errorf("test error"))
  assert.Equal(t, http.StatusNotFound, w.Code)

  w = httptest.NewRecorder()
  WriteError(w, http.StatusInternalServerError, fmt.Errorf("test error"))
  assert.Equal(t, http.StatusInternalServerError, w.Code)

  w = httptest.NewRecorder()
  WriteError(w, http.StatusBadRequest, fmt.Errorf("test error"))
  assert.Equal(t, http.StatusBadRequest, w.Code)

  w = httptest.NewRecorder()
  assert.Panics(t, func() { WriteError(w, http.StatusBadRequest, nil) })
}
