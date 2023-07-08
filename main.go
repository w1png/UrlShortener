package main

import (
	"log"
	"reflect"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/w1png/urlshortener/graphqlHandlers"
	"github.com/w1png/urlshortener/handlers"
	"github.com/w1png/urlshortener/logger"
	"github.com/w1png/urlshortener/middleware"
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
  err := utils.ConfigInstance.Init()
  if err != nil {
    return err
  }

  err = logger.InitLogger()
  if err != nil {
    return err
  }

	err = storage.InitSelectedStorage()
	if err != nil {
		return err
	}

  err = utils.InitGRPCConnection()
  if err != nil {
    return err
  }

	return nil
}

func main() {
	err := onStartup()
	if err != nil {
    log.Fatal(err.Error())
	}

	if reflect.TypeOf(storage.SelectedStorage) == reflect.TypeOf(&storage.PostgresStorage{}) {
		err = autoMigrate()
		if err != nil {
      logger.LoggerInstance.Fatal(err.Error())
		}
	}

  log.Println("Starting http server on port 8081")
	server := NewApiServer(":8081")
  
  server.UseMiddleware(middleware.LoggingMiddleware)
  server.UseMiddleware(middleware.PrometheusDurationMiddleware)
  server.UseMiddleware(middleware.PrometheusCounterMiddleware)

  server.RegisterHandlerFunc("/api/v1/urls", handlers.CreateUrl, "POST")
  server.RegisterHandlerFunc("/api/v1/urls/{alias}", handlers.GetUrl, "GET")
  server.RegisterHandlerFunc("/metrics", promhttp.Handler().ServeHTTP, "GET")

  graphQLHTTPHandler, err := graphqlHandlers.GetHTTPHandler()
  if err != nil {
    logger.LoggerInstance.Fatal(err.Error())
  }
  server.RegisterHandlerFunc("/graphql", graphQLHTTPHandler.ServeHTTP, "POST", "GET")
  // get url: curl -g 'http://localhost:8081/graphql?query={url(alias:"test"){alias,url}}'
  // create url: curl -g 'http://localhost:8081/graphql?query=mutation+_{createUrl(url:"test"){alias,url}}'

	err = server.Run()
	if err != nil {
    logger.LoggerInstance.Fatal(err.Error())
	}
}
