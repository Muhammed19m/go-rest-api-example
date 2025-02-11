package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"rest-api/internal/app"
	"rest-api/internal/config"
)

 


func main() {
	
	config, err := config.LoadConfig()
	if  err != nil {
		log.Fatal("Error import config.env: ", err)
	}


	ctx, cancel := context.WithCancel(context.Background())

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	
	go func() {
		<- signals 
		log.Println("program completion...")
		cancel()
	}()

	if err := app.Run(ctx, config); err != nil {
		log.Fatal(err)
	}
	
	
	
}


