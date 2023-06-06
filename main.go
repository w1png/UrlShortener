package main

import (
	"log"
	"os"

	"github.com/w1png/ozontest/models"
	"github.com/w1png/ozontest/utils"
)

func autoMigrate() error {
	db := utils.DB

	err := db.AutoMigrate(&models.Url{})
	if err != nil {
		return err
	}

	return nil
}

func onStartup() error {
	err := utils.InitDB()
	if err != nil {
		return err
	}

	return nil
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}

func main() {
	err := onStartup()
	if err != nil {
		log.Fatal(err)
	}

	err = autoMigrate()
	if err != nil {
		log.Fatal(err)
	}

  server := NewApiServer(getPort())
  err = server.Run()
  if err != nil {
    log.Fatal(err)
  }
}
