package utils

import (
	"os"
	"strconv"
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

  RedisHost string
  RedisPort string
  RedisPassword string
  RedisDatabase int
  RedisTestDatabase int

  UrlGRPCServiceHost string

  LoggerType string
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

  c.RedisHost = os.Getenv("REDIS_HOST")
  c.RedisPort = os.Getenv("REDIS_PORT")
  c.RedisPassword = os.Getenv("REDIS_PASSWORD")
  redisDatabase, ok := os.LookupEnv("REDIS_DATABASE")
  if ok {
    redisDatabaseInt, err := strconv.Atoi(redisDatabase)
    if err != nil {
      return NewEnvironmentVariableError("REDIS_DATABASE is not a number")
    }
    c.RedisDatabase = redisDatabaseInt
  }
  redisTestDatabase, ok := os.LookupEnv("REDIS_TEST_DATABASE")
  if ok {
    redisTestDatabaseInt, err := strconv.Atoi(redisTestDatabase)
    if err != nil {
      return NewEnvironmentVariableError("REDIS_TEST_DATABASE is not a number")
    }
    c.RedisTestDatabase = redisTestDatabaseInt
  }

  c.UrlGRPCServiceHost, ok = os.LookupEnv("URL_GRPC_SERVICE_HOST")
  if !ok {
    return NewEnvironmentVariableError("URL_GRPC_SERVICE_HOST")
  }

  c.LoggerType = os.Getenv("LOGGER_TYPE")
  
  return nil
}

