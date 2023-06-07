package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/w1png/ozontask/models"
	"github.com/w1png/ozontask/utils"
)

func CreateUrl(w http.ResponseWriter, r *http.Request) {
	type RequestBody struct {
		Url string `json:"url"`
	}

	var body RequestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid request body"))
		return
	}

  if body.Url == "" {
    utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid request body"))
    return
  }

	url, err := models.NewUrl(body.Url)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = url.Save()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

  type ResponseBody struct {
    Url string `json:"url"`
    Alias string `json:"alias"`
  }
	utils.WriteResponse(w, http.StatusCreated, ResponseBody{Url: url.Url, Alias: url.Alias})
}

func GetUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	alias := vars["alias"]

	url, err := models.GetUrlByAlias(alias)
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
