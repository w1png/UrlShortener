package main

import (
	"log"
	"os"
	"reflect"

	"github.com/w1png/urlshortener/handlers"
	"github.com/w1png/urlshortener/models"
	"github.com/w1png/urlshortener/storage"
	"github.com/w1png/urlshortener/utils"
)

func autoMigrate() error {
	db := storage.SelectedStorage.(*storage.PostgresStorage).DB

	err := db.AutoMigrate(&models.Url{})
	if err != nil {
		return err
	}

	return nil
}

func onStartup() error {
	err := storage.InitSelectedStorage()
	if err != nil {
		return err
	}

  grpc_host, ok := os.LookupEnv("GRPC_HOST")
  if !ok {
    return storage.NewEnvironmentVariableError("GRPC_HOST")
  }

  err = utils.InitGRPCConnection(grpc_host)
  if err != nil {
    return err
  }

	return nil
}

func main() {
	err := onStartup()
	if err != nil {
		log.Fatal(err)
	}

	if reflect.TypeOf(storage.SelectedStorage) == reflect.TypeOf(&storage.PostgresStorage{}) {
		err = autoMigrate()
		if err != nil {
			log.Fatal(err)
		}
	}

  log.Println("Starting http server on port 8081")
	server := NewApiServer(":8081")
  
  server.RegisterHandlerFunc("/api/v1/urls", handlers.CreateUrl, "POST")
  server.RegisterHandlerFunc("/api/v1/urls/{alias}", handlers.GetUrl, "GET")

	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
