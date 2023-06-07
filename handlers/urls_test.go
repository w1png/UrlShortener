package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"github.com/w1png/ozontest/models"
	"github.com/w1png/ozontest/utils"
)

func TestUrl(t *testing.T) {
  var testValidUrl models.Url

  utils.InitDB(true)
  utils.DB.AutoMigrate(&models.Url{})

  url := []byte(`{"url": "http://www.google.com"}`)
  req, err := http.NewRequest("POST", "/urls", bytes.NewBuffer(url))

  assert.Nil(t, err)

  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(CreateUrl)
  handler.ServeHTTP(rr, req)

  assert.Equal(t, http.StatusCreated, rr.Code)
  assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

  err = json.Unmarshal(rr.Body.Bytes(), &testValidUrl)
  assert.Nil(t, err)
  assert.NotNil(t, testValidUrl.Url)
  assert.NotNil(t, testValidUrl.Alias)


  req, err = http.NewRequest("GET", "/urls/" + testValidUrl.Alias, nil)
  assert.Nil(t, err)
  req = mux.SetURLVars(req, map[string]string{"alias": testValidUrl.Alias})

  rr = httptest.NewRecorder()
  handler = http.HandlerFunc(GetUrl)
  handler.ServeHTTP(rr, req)

  assert.Equal(t, http.StatusOK, rr.Code)
  assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

  var testValidUrl2 models.Url
  err = json.Unmarshal(rr.Body.Bytes(), &testValidUrl2)
  assert.Nil(t, err)

  assert.Equal(t, testValidUrl.Url, testValidUrl2.Url)


  req, err = http.NewRequest("GET", "/urls/thisisnotanalias", nil)
  assert.Nil(t, err)
  req = mux.SetURLVars(req, map[string]string{"alias": "thisisnotanalias"})
  
  rr = httptest.NewRecorder()
  handler = http.HandlerFunc(GetUrl)
  handler.ServeHTTP(rr, req)

  assert.Equal(t, http.StatusNotFound, rr.Code)
  assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
  assert.Equal(t, "{\"error\":\"Url not found\"}\n", rr.Body.String())
}
