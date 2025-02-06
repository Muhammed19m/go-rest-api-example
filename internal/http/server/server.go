package server

import (
	"fmt"
	"net/http"
	"rest-api/internal/database"
	"rest-api/internal/handler"

	"github.com/gorilla/mux"
)
 




func Run(db *database.Database, config Config) error {
	if err := config.Validate(); err != nil {
		return fmt.Errorf("validation server config: %w", err)
	}

	server := handler.Server{Database: db}

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/wallet", server.HandleTransaction).Methods("POST")
	router.HandleFunc("/api/v1/wallets/{WALLET_UUID}", server.HandleGetBalance).Methods("GET")
	http.Handle("/", router)

	err := http.ListenAndServe(fmt.Sprint(":", config.Port), nil)
	if err != nil {
		return fmt.Errorf("http listen and serve: %w", err)
	}

	return nil
}



type Config struct {
	Port int
}

func (c Config) Validate() error {
	if c.Port < 1 || c.Port > 65535 {
		return fmt.Errorf("port must be in the range 1–65535")	
	}	
	return nil
}