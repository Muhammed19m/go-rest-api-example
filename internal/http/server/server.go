package server

import (
	"errors"
	"fmt"
	"net/http"
	"rest-api/internal/database"
	"rest-api/internal/handler"

	"github.com/gorilla/mux"
)
 

type Server struct {
	http.Server
	db *database.Database
}

func Init(db *database.Database, config Config) (*Server, error) {
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("validation server config: %w", err)
	}

	server := Server {http.Server{ Addr: fmt.Sprint(":", config.Port)}, db}

	handler := handler.Handler{Database: db}

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/wallet", handler.HandleTransaction).Methods("POST")
	router.HandleFunc("/api/v1/wallets/{WALLET_UUID}", handler.HandleGetBalance).Methods("GET")
	http.Handle("/", router)

	return &server, nil
}


func (server *Server)Run() error {
	
	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("http listen and serve: %w", err)
	}
	
	return nil
}



var ErrInvalidPort = errors.New("port must be in the range 1–65535")	


type Config struct {
	Port int
}

func (c Config) Validate() error {
	if c.Port < 1 || c.Port > 65535 {
		return ErrInvalidPort 
	}	
	return nil
}