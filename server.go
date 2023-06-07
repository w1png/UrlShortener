package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/w1png/ozontask/handlers"
)

type ApiServer struct {
	listenAddr string
}

func NewApiServer(port string) *ApiServer {
	return &ApiServer{
		listenAddr: port,
	}
}

func (s *ApiServer) Run() error {
	router := mux.NewRouter()

	router.HandleFunc("/urls", handlers.CreateUrl).Methods("POST")
	router.HandleFunc("/urls/{alias}", handlers.GetUrl).Methods("GET")

	log.Printf("Listening on %s\n", s.listenAddr)
	return http.ListenAndServe(s.listenAddr, router)
}
