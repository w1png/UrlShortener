package main

import (
	"log"
	"net"

	"github.com/w1png/urlshortener/pkg/url"
	"github.com/w1png/urlshortener/pkg/url/endpoints"
	"github.com/w1png/urlshortener/pkg/url/transport"
	"github.com/w1png/urlshortener/storage"
	"google.golang.org/grpc"

	pb "github.com/w1png/urlshortener/pkg/url/proto"
)

const port = ":8080"

func main() {
  err := storage.InitSelectedStorage()
  if err != nil {
    log.Fatal(err)
  }

  var (
    svc = url.NewUrlService()
    endpoints = endpoints.NewSet(svc)
    grpcServer = transport.NewGRPCServer(endpoints)
  )

  listener, err := net.Listen("tcp", port)
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }

  server := grpc.NewServer()
  pb.RegisterUrlServiceServer(server, grpcServer)

  log.Printf("Starting gRPC server on port %s", port)
  server.Serve(listener)
}
