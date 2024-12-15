package dbhand

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"rest-api/internal/model"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	var err error
	url_db := os.Getenv("DATABASE_URL")
	db, err = sql.Open("postgres", url_db)
	if err != nil {
		log.Fatal("Error conection: ", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}



func GetBalanceByUUID(uuid int) (int, error) {
	var balance int
	err := db.QueryRow("SELECT balance FROM wallet WHERE wallet_id = $1;", uuid).Scan(&balance)
	return balance, err
}


func CreateWalletByUUID(uuid int, amount int) error {
	_, err:= db.Query("INSERT INTO wallet VALUES ($1, $2);", uuid, 0)
	return err
}

func Deposit(transaction model.Transaction) error {
	balance, err := GetBalanceByUUID(transaction.WalletId)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			CreateWalletByUUID(transaction.WalletId, transaction.Amount)
			return nil
 		}
		return err
	}
	_, err = db.Query("UPDATE wallet SET balance = $1 WHERE wallet_id = $2;", balance+transaction.Amount, transaction.WalletId)
	return err
}

func Withdraw(transaction model.Transaction) error { 
	balance, err := GetBalanceByUUID(transaction.WalletId)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			CreateWalletByUUID(transaction.WalletId, transaction.Amount)
			return errors.New("not enough money")
 		}
		return err
	}

	if balance < transaction.Amount {
		return errors.New("not enough money")
	} else {
		_, err = db.Query("UPDATE wallet SET balance = $1 WHERE wallet_id = $2;", balance-transaction.Amount, transaction.WalletId)
		return err
	}
}