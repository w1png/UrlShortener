package transport

import (
	"context"
	"fmt"

	"github.com/w1png/urlshortener/pkg/url/endpoints"
	pb "github.com/w1png/urlshortener/pkg/url/proto"

	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
  pb.UnsafeUrlServiceServer

  createUrlHandler grpctransport.Handler
  getUrlHandler grpctransport.Handler
}

func NewGRPCServer(endpoints endpoints.Set) pb.UrlServiceServer {
  return &grpcServer{
    createUrlHandler: grpctransport.NewServer(
      endpoints.CreateUrlEndpoint,
      decodeGRPCCreateUrlRequest,
      encodeGRPCCreateUrlResponse,
    ),
    getUrlHandler: grpctransport.NewServer(
      endpoints.GetUrlEndpoint,
      decodeGRPCGetUrlRequest,
      encodeGRPCGetUrlResponse,
    ),
  }
}

func (s *grpcServer) CreateUrl(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
  fmt.Println("CreateUrl")
  _, resp, err := s.createUrlHandler.ServeGRPC(ctx, req)
  if err != nil {
    return nil, err
  }
  return resp.(*pb.CreateResponse), nil
}

func (s *grpcServer) GetUrl(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
  _, resp, err := s.getUrlHandler.ServeGRPC(ctx, req)
  if err != nil {
    return nil, err
  }
  return resp.(*pb.GetResponse), nil
}

func decodeGRPCCreateUrlRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
  fmt.Println("decodeGRPCCreateUrlRequest")
  req := grpcReq.(*pb.CreateRequest)
  return endpoints.CreateUrlRequest{Url: req.Url}, nil
}

func encodeGRPCCreateUrlResponse(_ context.Context, response interface{}) (interface{}, error) {
  fmt.Println("encodeGRPCCreateUrlResponse")
  resp := response.(endpoints.CreateUrlResponse)
  return &pb.CreateResponse{Url: resp.Url, Alias: resp.Alias}, nil
}

func decodeGRPCGetUrlRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
  req := grpcReq.(*pb.GetRequest)
  return endpoints.GetUrlRequest{Alias: req.Alias}, nil
}

func encodeGRPCGetUrlResponse(_ context.Context, response interface{}) (interface{}, error) {
  resp := response.(endpoints.GetUrlResponse)
  return &pb.GetResponse{Url: resp.Url, Alias: resp.Alias}, nil
}

