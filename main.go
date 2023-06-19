package main

import (
	"log"
	"reflect"

	"github.com/w1png/urlshortener/models"
	"github.com/w1png/urlshortener/storage"
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

	server := NewApiServer(":8081")
	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
