package server

import (
	"fmt"
	"net/http"
	"rest-api/internal/handler"

	"github.com/gorilla/mux"
)


 

func Run(config Config) error {
	if err := config.Validate(); err != nil {
		return fmt.Errorf("validation server config: %w", err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/wallet", handler.HandleTransaction).Methods("POST")
	router.HandleFunc("/api/v1/wallets/{WALLET_UUID}", handler.GetBalance).Methods("GET")
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
	if c.Port >= 1 && c.Port <= 65535 {
		return nil
	}
	return fmt.Errorf("port must be in the range 1–65535")
}