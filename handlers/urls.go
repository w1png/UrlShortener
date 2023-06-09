package handlers

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
	"github.com/w1png/urlshortener/models"
	"github.com/w1png/urlshortener/storage"
	"github.com/w1png/urlshortener/utils"
)

func CreateUrl(w http.ResponseWriter, r *http.Request) {
	type RequestBody struct {
		Url string `json:"url"`
	}

	var body RequestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, NewInvalidRequestBodyError(err.Error()))
		return
	}

  if body.Url == "" {
    utils.WriteError(w, http.StatusBadRequest, NewInvalidRequestBodyError("Url is empty"))
    return
  }

  url := models.NewUrl(body.Url)
  _, err = storage.SelectedStorage.GetByAlias(url.Alias)
  if err == nil && url != nil {
    utils.WriteError(w, http.StatusInternalServerError, NewUniqueAliasError("Failed to create unique alias"))
    return
  }
  if err != nil && reflect.TypeOf(err) != reflect.TypeOf(&storage.NotFoundError{}) {
    utils.WriteError(w, http.StatusInternalServerError, err)
    return
  }

	err = storage.SelectedStorage.Save(url)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

  type ResponseBody struct {
    Url string `json:"url"`
    Alias string `json:"alias"`
  }
	utils.WriteResponse(w, http.StatusOK, ResponseBody{Url: url.Url, Alias: url.Alias})
}

func GetUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	alias := vars["alias"]

	url, err := storage.SelectedStorage.GetByAlias(alias)
  if err != nil {
    utils.WriteError(w, http.StatusNotFound, err)
    return
  }

  if url.Url == "" || url.Alias == "" {
    utils.WriteError(w, http.StatusNotFound, err)
    return
  }

  type ResponseBody struct {
    Url string `json:"url"`
    Alias string `json:"alias"`
  }
	utils.WriteResponse(w, http.StatusOK, ResponseBody{Url: url.Url, Alias: url.Alias})
}
