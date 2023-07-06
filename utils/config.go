package utils

import (
	"os"
)

var ConfigInstance Config

type Config struct {
  StorageType string
  PostgresHost string
  PostgresPort string
  PostgresUser string
  PostgresPassword string
  PostgresDatabase string
  PostgresTestDatabase string

  UrlGRPCServiceHost string
}

func (c *Config) Init() error {
  storage_type, ok := os.LookupEnv("STORAGE_TYPE")
  if !ok {
    return NewEnvironmentVariableError("STORAGE_TYPE")
  }
  
  c.StorageType = storage_type
  c.PostgresHost = os.Getenv("POSTGRES_HOST")
  c.PostgresPort = os.Getenv("POSTGRES_PORT")
  c.PostgresUser = os.Getenv("POSTGRES_USER")
  c.PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
  c.PostgresDatabase = os.Getenv("POSTGRES_DATABASE")
  c.PostgresTestDatabase = os.Getenv("POSTGRES_TEST_DATABASE")

  c.UrlGRPCServiceHost = os.Getenv("URL_GRPC_SERVICE_HOST")
  
  return nil
}

