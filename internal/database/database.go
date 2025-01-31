package database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"rest-api/internal/model"

	_ "github.com/lib/pq"
)

// db - синглтон экземпляр базы данных
// todo: удалить синглтон и сделать экземпляр бд возращаемым
var db *sql.DB

var ErrConnection = errors.New("error connection")


// InitDB инициализирует соединение с бд.
// изменить сигнатуру функции:
// InitDB(config Config) (*sql.DB, error)
func InitDB() error{
	var err error
	url_db := os.Getenv("DATABASE_URL")
	db, err = sql.Open("postgres", url_db)
	if err != nil {
		return fmt.Errorf("sql open: %w", err)
	}
	
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("dp ping: %w", err)
	} 
	return nil
}



func GetBalanceByUUID(uuid int) (int, error) {
	var balance int
	err := db.QueryRow("SELECT balance FROM wallet WHERE wallet_id = $1;", uuid).Scan(&balance)
	return balance, err

}


func CreateWalletByUUID(uuid int, amount int) error {
	_, err:= db.Exec("INSERT INTO wallet VALUES ($1, $2);", uuid, amount)
	return err
}


func Deposit(transaction model.Transaction) error {
	balance, err := GetBalanceByUUID(transaction.WalletId)
	if err != nil {
		// todo: использовать errors.Is
		if err.Error() == "sql: no rows in result set" {
			CreateWalletByUUID(transaction.WalletId, transaction.Amount)
			return nil
 		}
		return err
	}
	_, err = db.Exec("UPDATE wallet SET balance = $1 WHERE wallet_id = $2;", balance+transaction.Amount, transaction.WalletId)
	//  todo: использовать update on conflict
	// db.Exec("INSERT into wallet(wallet_id, balance) VALUES ($1, $2)", balance+transaction.Amount, transaction.WalletId)
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

	// todo: правила бизнес логики проверять перед функциями database
	if balance < transaction.Amount {
		return errors.New("not enough money")
	} else {

		_, err = db.Exec("UPDATE wallet SET balance = $1 WHERE wallet_id = $2;", balance-transaction.Amount, transaction.WalletId)
		return err
	}
}