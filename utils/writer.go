package utils

import (
	"encoding/json"
	"net/http"
)

type Error interface {
  Error() string
}

func WriteResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func WriteError(w http.ResponseWriter, status int, err interface{}) {
  var errorText string

  switch err.(type) {
  case string:
    errorText = err.(string)
  case Error:
    errorText = err.(Error).Error()
  default:
    errorText = "Internal Server Error"
  }

	WriteResponse(w, status, map[string]string{
    "error": errorText,
  })
}
