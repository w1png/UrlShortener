package main

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"

	pb "github.com/w1png/urlshortener/pkg/url/proto"
	"github.com/w1png/urlshortener/utils"
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

  conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
  if err != nil {
    return err
  }
  defer conn.Close()

  router.HandleFunc("/api/v1/urls", func(w http.ResponseWriter, r *http.Request) {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()


    client := pb.NewUrlServiceClient(conn)
    resp, err := client.CreateUrl(ctx, &pb.CreateRequest{Url: "http://google.com"})
    if err != nil {
      utils.WriteError(w, http.StatusInternalServerError, err)
      return
    }
    
    utils.WriteResponse(w, http.StatusOK, resp)
  }).Methods("POST")
    
  return http.ListenAndServe(s.listenAddr, router)
}
