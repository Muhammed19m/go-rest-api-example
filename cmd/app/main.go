package main

import (
	"log"

	"github.com/joho/godotenv"

	"rest-api/internal/app"
)

 


func main() {
	// todo: узнать о graceful shutdown golang

	if err:= loadConfig(); err != nil {
		log.Fatal("Error import config.env: ", err)
	}
	if err := app.Run(/*  */); err != nil {
		log.Fatal(err)
	}
	
	log.Println("App finished")
}


// todo: пенести в пакет config, использовать структуру Config
func loadConfig() error {
	if err := godotenv.Load("config.env"); err != nil {
        return err
    }
    return nil
}

