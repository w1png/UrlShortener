package main

import (
	"context"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/w1png/urlshortener/pkg/url"
	"github.com/w1png/urlshortener/pkg/url/endpoints"
	pb "github.com/w1png/urlshortener/pkg/url/proto"
	"github.com/w1png/urlshortener/pkg/url/transport"
	"github.com/w1png/urlshortener/storage"
	"github.com/w1png/urlshortener/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

var listener *bufconn.Listener

func init() {
  err := utils.ConfigInstance.Init()
  if err != nil {
    panic(err)
  }
  utils.ConfigInstance.StorageType = "in_memory"
  err = storage.InitSelectedStorage()
  if err != nil {
    panic(err)
  }

  listener = bufconn.Listen(1024 * 1024)

  var (
    svc = url.NewUrlService()
    endpoints = endpoints.NewSet(svc)
    grpcServer = transport.NewGRPCServer(endpoints)
  )
  server := grpc.NewServer()
  pb.RegisterUrlServiceServer(server, grpcServer)

  go func() {
    if err := server.Serve(listener); err != nil {
      panic(err)
    }
  }()
}

func bufDialer(context.Context, string) (net.Conn, error) {
  return listener.Dial()
}

func TestCreateUrl(t *testing.T) {
  ctx := context.Background()
  conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
  assert.Nil(t, err)
  defer conn.Close()
  client := pb.NewUrlServiceClient(conn)

  req := &pb.CreateRequest{Url: "https://www.google.com"}
  res, err := client.CreateUrl(ctx, req)
  assert.Nil(t, err)

  assert.Equal(t, "https://www.google.com", res.Url)
  assert.NotEmpty(t, res.Alias)
}

func TestGetUrl(t *testing.T) {
  ctx := context.Background()
  conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
  assert.Nil(t, err)
  defer conn.Close()
  client := pb.NewUrlServiceClient(conn)

  req := &pb.CreateRequest{Url: "https://www.google.com"}
  res, err := client.CreateUrl(ctx, req)
  assert.Nil(t, err)

  req2 := &pb.GetRequest{Alias: res.Alias}
  res2, err := client.GetUrl(ctx, req2)
  assert.Nil(t, err)

  assert.Equal(t, "https://www.google.com", res2.Url)
  assert.Equal(t, res.Alias, res2.Alias)
}

