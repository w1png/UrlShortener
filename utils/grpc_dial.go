package utils

import "google.golang.org/grpc"

var GRPCConnection *grpc.ClientConn

func InitGRPCConnection(host string) error {
  var err error

  GRPCConnection, err = grpc.Dial(host, grpc.WithInsecure())
  if err != nil {
    return err
  }

  return nil
}
