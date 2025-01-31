package server

import (
	"fmt"
	"net/http"
	"rest-api/internal/handler"

	"github.com/gorilla/mux"
)

type Config struct {
	Port string
}

func Run(config Config) error {
	// config.validate()
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/wallet", handler.HandleTransaction).Methods("POST")
	router.HandleFunc("/api/v1/wallets/{WALLET_UUID}", handler.GetBalance).Methods("GET")
	http.Handle("/", router)

	// port := os.Getenv("PORT")
	err := http.ListenAndServe(":" + config.Port, nil)
	if err != nil {
		return fmt.Errorf("http listen and serve: %w", err)
	}

	return nil
}