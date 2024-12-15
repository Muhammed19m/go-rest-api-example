package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	dbhand "rest-api/internal/db"
	"rest-api/internal/handler"
)

 


func main() {
	if err:= loadConfig(); err != nil {
		log.Fatal("Error import config.env: ", err)
	}

	dbhand.InitDB()
	// defer dbhand.Close()
	
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/wallet", handler.HandleTransaction).Methods("POST")
	router.HandleFunc("/api/v1/wallets/{WALLET_UUID}", handler.GetBalance).Methods("GET")
	http.Handle("/", router)

	port := os.Getenv("PORT")
 	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		log.Fatal("Server startup error!")
	} 

}



func loadConfig() error {
	if err := godotenv.Load("config.env"); err != nil {
        return err
    }
    return nil
}

