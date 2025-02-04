package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"rest-api/internal/database"
	"rest-api/internal/model"
	"strconv"

	"github.com/gorilla/mux"
)



type Server struct {
	Database *sql.DB
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

	}
	// +валидация OperationType
	// +вызов бизнес логики с передачей параметров, полученные из запроса
	// +обработка результата
	if transaction.OperationType == model.DEPOSIT {
		if err := database.Deposit(s.Database, transaction); err != nil {
			http.Error(w, "error "+err.Error(), http.StatusInternalServerError)
		}
	} else if transaction.OperationType == model.WITHDRAW {
		
		if err:=database.Withdraw(s.Database, transaction); err != nil {
			http.Error(w, "error"+err.Error(), http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Invalid operation type", http.StatusBadRequest)
		return
	}
}
	
func UnmarshalBody(r *http.Request, a any) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return err
	}
	defer r.Body.Close()
	err = json.Unmarshal(body, a)
	if err != nil {
		// http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return err
	}

	return nil
}

	
	
		// функция обработчик, которая обрабатывает метод GET
		// принимает ID кошелька и возращает баланс 
		
		// запрос: `GET api/v1/wallets/{WALLET_UUID}`
func (s *Server) GetBalance(w http.ResponseWriter, r *http.Request) {
	wallet_uuid := mux.Vars(r)["WALLET_UUID"]

	uuid, _/* err_conv */ := strconv.ParseInt(wallet_uuid, 10, 32) 

	balance, err := database.GetBalanceByUUID(s.Database, int(uuid))

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			database.CreateWalletByUUID(s.Database, int(uuid), 0)
			w.Write([]byte("0"))
		} else {
			http.Error(w, "Invalid UUID format", http.StatusBadRequest)
		}
	} else {
		w.Write([]byte(fmt.Sprint(balance)))
	}
}

// func getIntVarFromPath(path string) (int, error) {}



	