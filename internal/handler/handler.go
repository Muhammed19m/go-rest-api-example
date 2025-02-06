package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"rest-api/internal/database"
	"rest-api/internal/model"
	"rest-api/internal/service"

	"github.com/gorilla/mux"
)



type Server struct {
	Database *database.Database
	// Service service.Service
}



// функция обработчик, которая обрабатывает метод POST
// запрос:
// `POST api/v1/wallet
// {
// 	valletId: UUID,
// 	operationType: DEPOSIT or WITHDRAW,
// 	amount: 1000
// }`
 func (s *Server) HandleTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction model.Transaction

	if err := UnmarshalBody(r, &transaction); err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := service.ProccesTransaction(s.Database, transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
 }


// функция обработчик, которая обрабатывает метод GET
// принимает ID кошелька и возращает баланс 
		
// запрос: `GET api/v1/wallets/{WALLET_UUID}`
 func (s *Server) HandleGetBalance(w http.ResponseWriter, r *http.Request) {
	wallet_uuid := mux.Vars(r)["WALLET_UUID"]
	uuid, err := strconv.ParseInt(wallet_uuid, 10, 32) 

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	balance, err := service.GetBalance(s.Database, int(uuid))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp := ResponseBalance {balance}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
 	json.NewEncoder(w).Encode(resp)

 }


	
type ResponseBalance struct {
	Balance int `json:"balance"`
}