package main

import (
	"net/http"

	"github.com/gorilla/handlers"
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
  return http.ListenAndServe(s.listenAddr, handlers.CORS(
    handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
    handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
    handlers.AllowedOrigins([]string{"*"}),
  )(s.router))
}

func (s *ApiServer) RegisterHandlerFunc(path string, f func(w http.ResponseWriter, r *http.Request), methods ...string) {
  s.router.HandleFunc(path, f).Methods(methods...)
}

func (s *ApiServer) UseMiddleware(f func(http.Handler) http.Handler) {
  s.router.Use(f)
}

