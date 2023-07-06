package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	listenAddr string
  router *mux.Router
}

func NewApiServer(port string) *ApiServer {
	return &ApiServer{
		listenAddr: port,
    router: mux.NewRouter(),
	}
}

func (s *ApiServer) Run() error {
  return http.ListenAndServe(s.listenAddr, s.router)
}

func (s *ApiServer) RegisterHandlerFunc(path string, f func(w http.ResponseWriter, r *http.Request), method string) {
  s.router.HandleFunc(path, f).Methods(method)
}

func (s *ApiServer) UseMiddleware(f func(http.Handler) http.Handler) {
  s.router.Use(f)
}

