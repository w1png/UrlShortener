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
	if os.Getenv("USE_IN_MEMORY") == "true" {
		utils.UseIM = true
		return nil
	}

	err := utils.InitDB(false)
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

	if !utils.UseIM {
		err = autoMigrate()
		if err != nil {
			log.Fatal(err)
		}
	}

	server := NewApiServer(":8080")
	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
