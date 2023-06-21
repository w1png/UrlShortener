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

func TestWriteResponse_StatusOK(t *testing.T) {
  w := httptest.NewRecorder()
  WriteResponse(w, http.StatusOK, "test")
  assert.Equal(t, http.StatusOK, w.Code)
}

func TestWriteResponse_StatusNotFound(t *testing.T) {
  w := httptest.NewRecorder()
  WriteResponse(w, http.StatusNotFound, nil)
  assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestWriteResponse_StatusInternalServerError(t *testing.T) {
  w := httptest.NewRecorder()
  WriteResponse(w, http.StatusInternalServerError, nil)
  assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestWriteResponse_StatusBadRequest(t *testing.T) {
  w := httptest.NewRecorder()
  WriteResponse(w, http.StatusBadRequest, nil)
  assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestWriteError_String(t *testing.T) {
  w := httptest.NewRecorder()
  WriteError(w, http.StatusNotFound, "test error")
  assert.Equal(t, http.StatusNotFound, w.Code)
  assert.Equal(t, "{\"error\":\"test error\"}\n", w.Body.String())
}

func TestWriteError_Error(t *testing.T) {
  w := httptest.NewRecorder()
  WriteError(w, http.StatusBadRequest, fmt.Errorf("test error"))
  assert.Equal(t, http.StatusBadRequest, w.Code)
  assert.Equal(t, "{\"error\":\"test error\"}\n", w.Body.String())
}

func TestWriteError_Default(t *testing.T) {
  w := httptest.NewRecorder()
  WriteError(w, http.StatusBadRequest, TestObject{"test", "test"})
  assert.Equal(t, http.StatusBadRequest, w.Code)
  assert.Equal(t, "{\"error\":\"Internal Server Error\"}\n", w.Body.String())
}
