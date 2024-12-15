package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	dbhand "rest-api/internal/db"
	"rest-api/internal/model"
	"strconv"

	"github.com/gorilla/mux"
)

// функция обработчик, которая обрабатывает метод POST
// запрос:
// `POST api/v1/wallet
// {
// 	valletId: UUID,
// 	operationType: DEPOSIT or WITHDRAW,
// 	amount: 1000
// }`
func HandleTransaction(w http.ResponseWriter, r *http.Request) {
	
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	var transaction model.Transaction
	err = json.Unmarshal(body, &transaction)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	if transaction.OperationType == model.DEPOSIT {
		if err:=dbhand.Deposit(transaction); err != nil {
			http.Error(w, "error "+err.Error(), http.StatusInternalServerError)
		}
	} else if transaction.OperationType == model.WITHDRAW {
		
		if err:=dbhand.Withdraw(transaction); err != nil {
			http.Error(w, "error"+err.Error(), http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Invalid operation type", http.StatusBadRequest)
		return
	}
}
	
	
	
		// функция обработчик, которая обрабатывает метод GET
		// принимает ID кошелька и возращает баланс 
		
		// запрос: `GET api/v1/wallets/{WALLET_UUID}`
func GetBalance(w http.ResponseWriter, r *http.Request) {
	wallet_uuid := mux.Vars(r)["WALLET_UUID"]

	uuid, _/* err_conv */ := strconv.ParseInt(wallet_uuid, 10, 32) 

	balance, err := dbhand.GetBalanceByUUID(int(uuid))

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			dbhand.CreateWalletByUUID(int(uuid), 0)
			w.Write([]byte("0"))
		} else {
			http.Error(w, "Invalid UUID format", http.StatusBadRequest)
		}
	} else {
		w.Write([]byte(fmt.Sprint(balance)))
	}
}



	