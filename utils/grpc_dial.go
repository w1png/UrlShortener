package utils

import (
	"google.golang.org/grpc"
)

var UrlGRPCConnection *grpc.ClientConn

func InitGRPCConnection() error {
  var err error

  UrlGRPCConnection, err = grpc.Dial(ConfigInstance.UrlGRPCServiceHost, grpc.WithInsecure())
  if err != nil {
    return err
  }

  return nil
}
