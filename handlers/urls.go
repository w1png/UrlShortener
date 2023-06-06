package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/w1png/ozontest/models"
	"github.com/w1png/ozontest/utils"
)

func CreateUrl(w http.ResponseWriter, r *http.Request) {
	type RequestBody struct {
		Url string `json:"url"`
	}

	var body RequestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
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

	utils.WriteResponse(w, http.StatusCreated, url.Alias)
}

func GetUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	alias := vars["alias"]

	url, err := models.GetUrlByAlias(alias)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	utils.WriteResponse(w, http.StatusOK, url.Url)
}
