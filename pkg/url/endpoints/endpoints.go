package endpoints

import (
	"context"
	"net/http"
	"reflect"

	"github.com/go-kit/kit/endpoint"
	"github.com/w1png/urlshortener/pkg/url"
	"github.com/w1png/urlshortener/storage"
	"google.golang.org/grpc/status"
)

type Set struct {
  CreateUrlEndpoint endpoint.Endpoint
  GetUrlEndpoint endpoint.Endpoint
}

func NewSet(svc url.Service) Set {
  return Set{
    CreateUrlEndpoint: MakeCreateUrlEndpoint(svc),
    GetUrlEndpoint: MakeGetUrlEndpoint(svc),
  }
}

func MakeCreateUrlEndpoint(svc url.Service) endpoint.Endpoint {
  return func(ctx context.Context, request interface{}) (interface{}, error) {
    req := request.(CreateUrlRequest)
    url, err := svc.CreateUrl(req.Url)
    if err != nil {
      return CreateUrlResponse{}, status.Error(http.StatusInternalServerError, err.Error())
    }
    return CreateUrlResponse{url.Url, url.Alias}, nil
  }
}

func MakeGetUrlEndpoint(svc url.Service) endpoint.Endpoint {
  return func(ctx context.Context, request interface{}) (interface{}, error) {
    req := request.(GetUrlRequest)
    url, err := svc.GetUrl(req.Alias)
    if err != nil {
      if reflect.TypeOf(err) == reflect.TypeOf(&storage.NotFoundError{}) {
        return CreateUrlResponse{}, status.Error(http.StatusNotFound, err.Error())
      }
      return GetUrlResponse{}, status.Error(http.StatusInternalServerError, err.Error())
    }
    return GetUrlResponse{url.Url, url.Alias}, nil
  }
}

