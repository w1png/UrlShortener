package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/w1png/urlshortener/utils"
	"google.golang.org/grpc/status"

	endpoints "github.com/w1png/urlshortener/pkg/url/endpoints"
	pb "github.com/w1png/urlshortener/pkg/url/proto"
)

func GetUrl(w http.ResponseWriter, r *http.Request) {
  alias := mux.Vars(r)["alias"]

  if alias == "" {
    utils.WriteError(w, http.StatusBadRequest, utils.NewInvalidRequestBodyError("alias"))
    return
  }

  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()

  client := pb.NewUrlServiceClient(utils.UrlGRPCConnection)
  resp, err := client.GetUrl(ctx, &pb.GetRequest{Alias: alias})
  if err != nil {
    grpcErr, ok := status.FromError(err)
    if ok {
      utils.WriteError(w, int(grpcErr.Code()), grpcErr.Message())
      return
    }

    utils.WriteError(w, http.StatusInternalServerError, err)
    return
  }

  utils.WriteResponse(w, http.StatusOK, resp)
}

func CreateUrl(w http.ResponseWriter, r *http.Request) {
  var req endpoints.CreateUrlRequest
  err := json.NewDecoder(r.Body).Decode(&req)
  if err != nil {
    utils.WriteError(w, http.StatusBadRequest, utils.NewInvalidRequestBodyError("url"))
    return
  }

  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()

  client := pb.NewUrlServiceClient(utils.UrlGRPCConnection)
  resp, err := client.CreateUrl(ctx, &pb.CreateRequest{Url: req.Url})
  if err != nil {
    grpcErr, ok := status.FromError(err)
    if ok {
      utils.WriteError(w, int(grpcErr.Code()), grpcErr.Message())
      return
    }

    utils.WriteError(w, http.StatusInternalServerError, err)
    return
  }
  
  utils.WriteResponse(w, http.StatusOK, resp)
}

