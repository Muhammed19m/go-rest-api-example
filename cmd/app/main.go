package main

import (
	"log"

	"rest-api/internal/app"
	"rest-api/internal/config"
)

 


func main() {
	// todo: узнать о graceful shutdown golang
	
	config, err := config.LoadConfig()

	if  err != nil {
		log.Fatal("Error import config.env: ", err)
	}
	if err := app.Run(config); err != nil {
		log.Fatal(err)
	}
	
	log.Println("App finished")
}


