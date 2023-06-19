package endpoints

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
	"github.com/w1png/urlshortener/pkg/url"
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
    fmt.Println("MakeCreateUrlEndpoint")
    req := request.(CreateUrlRequest)
    url, err := svc.CreateUrl(req.Url)
    if err != nil {
      return CreateUrlResponse{}, err
    }
    return CreateUrlResponse{url.Url, url.Alias}, nil
  }
}

func MakeGetUrlEndpoint(svc url.Service) endpoint.Endpoint {
  return func(ctx context.Context, request interface{}) (interface{}, error) {
    req := request.(GetUrlRequest)
    url, err := svc.GetUrl(req.Alias)
    if err != nil {
      return GetUrlResponse{url.Url, url.Alias}, nil
    }
    return nil, err
  }
}

