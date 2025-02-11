package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"rest-api/internal/config"
	"rest-api/internal/database"
	"rest-api/internal/http/server"
)

const (TimeOutShutdownApp = time.Second * 5)


func Run(ctx context.Context, config *config.Config) error {
	
	db, err := database.InitDB(config)
	
	if err != nil {
		return fmt.Errorf("app init database: %w", err)
	}
	defer db.Close()


	server, err := server.Init(db, server.Config {Port: config.Port()})
	if err != nil {
		return fmt.Errorf("server init: %w", err)
	}

	var wg sync.WaitGroup
	defer wg.Wait()
	wg.Add(1)
	go func (wg *sync.WaitGroup) {
		<- ctx.Done()
		db.Close()
		ctxserv, cancel := context.WithTimeout(context.Background(), TimeOutShutdownApp)
		defer cancel()
		if err := server.Shutdown(ctxserv); err != nil {
			log.Println("graceful shutdown: ", err)
		} else {
			log.Println("server terminated")
		}
		wg.Done()
	} (&wg)


	if err:= server.Run(); errors.Is(err, http.ErrServerClosed) {
	} else if  err != nil {
		return fmt.Errorf("server run: %w", err)
	}
	
	
	return nil	
}