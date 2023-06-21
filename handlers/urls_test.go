package handlers
//
// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"regexp"
// 	"testing"
//
// 	"github.com/gorilla/mux"
// 	"github.com/stretchr/testify/assert"
//
// 	"github.com/w1png/urlshortener/models"
// 	"github.com/w1png/urlshortener/storage"
// )
//
// type ErrorResponse struct {
//   Error string `json:"error"`
// }
//
// func setup(t *testing.T) {
//   var err error
//   storage.SelectedStorage, err = storage.NewPostgresStorage(true)
//   if err != nil {
//     t.Fatal(err)
//   }
// }
//
//
// func TestCreateUrl(t *testing.T) {
//   setup(t)
//
//   req, err := http.NewRequest("POST", "/urls", bytes.NewBuffer([]byte(`{"url":"https://google.com"}`)))
//   if err != nil {
//     t.Fatal(err)
//   }
//
//   rr := httptest.NewRecorder()
//   handler := http.HandlerFunc(CreateUrl)
//   handler.ServeHTTP(rr, req)
//
//   assert.Equal(t, http.StatusOK, rr.Code)
//   assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
//   
//   var response models.Url
//   err = json.Unmarshal(rr.Body.Bytes(), &response)
//   assert.Nil(t, err)
//
//   assert.Equal(t, "https://google.com", response.Url)
//   assert.Equal(t, 10, len(response.Alias))
//
//   match, err := regexp.MatchString("^[a-zA-Z0-9_]{10}$", response.Alias)
//   assert.Nil(t, err)
//   assert.True(t, match)
// }
//
// func TestCreateUrl_InvalidRequestBodyError(t *testing.T) {
//   setup(t)
//
//   var reqs []http.Request
//
//   for _, body := range []string{`{"url:"https://google.com"`, `{"invalidkey":"https://google.com"}`} {
//     req, err := http.NewRequest("POST", "/urls", bytes.NewBuffer([]byte(body)))
//     if err != nil {
//       t.Fatal(err)
//     }
//     reqs = append(reqs, *req)
//   }
//
//   for _, req := range reqs {
//     t.Logf("Testing request body: %s", req.Body)
//
//     rr := httptest.NewRecorder()
//     handler := http.HandlerFunc(CreateUrl)
//     handler.ServeHTTP(rr, &req)
//
//     assert.Equal(t, http.StatusBadRequest, rr.Code)
//     assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
//
//     var response ErrorResponse
//     err := json.Unmarshal(rr.Body.Bytes(), &response)
//     assert.Nil(t, err)
//
//     assert.NotNil(t, response.Error)
//   }
// }
//
// func TestGetUrl(t *testing.T) {
//   setup(t)
//
//   url := models.NewUrl("https://google.com")
//   err := storage.SelectedStorage.Save(url)
//   if err != nil {
//     t.Fatal(err)
//   }
//
//   req, err := http.NewRequest("GET", "/urls/" + url.Alias, nil)
//   if err != nil {
//     t.Fatal(err)
//   }
//   req = mux.SetURLVars(req, map[string]string{"alias": url.Alias})
//
//   rr := httptest.NewRecorder()
//   handler := http.HandlerFunc(GetUrl)
//   handler.ServeHTTP(rr, req)
//
//   assert.Equal(t, http.StatusOK, rr.Code)
//   assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
//
//   var url2 models.Url
//   err = json.Unmarshal(rr.Body.Bytes(), &url2)
//   assert.Nil(t, err)
//
//   assert.Equal(t, url.Url, url2.Url)
//   assert.Equal(t, url.Alias, url2.Alias)
// }
//
// func TestGetUrl_NotFound(t *testing.T) {
//   setup(t)
//
//   req, err := http.NewRequest("GET", "/urls/invalidalias", nil)
//   if err != nil {
//     t.Fatal(err)
//   }
//   req = mux.SetURLVars(req, map[string]string{"alias": "invalidalias"})
//
//   rr := httptest.NewRecorder()
//   handler := http.HandlerFunc(GetUrl)
//   handler.ServeHTTP(rr, req)
//
//   assert.Equal(t, http.StatusNotFound, rr.Code)
//   assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
//
//   var response ErrorResponse
//   err = json.Unmarshal(rr.Body.Bytes(), &response)
//   assert.Nil(t, err)
//
//   assert.NotNil(t, response.Error)
// }
//
