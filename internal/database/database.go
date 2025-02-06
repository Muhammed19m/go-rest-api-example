package database

import (
	"database/sql"
	"errors"
	"fmt"
	"rest-api/internal/config"
	"rest-api/internal/model"

	_ "github.com/lib/pq"
)


var ErrConnection = errors.New("error connection")

type Database struct {
	db *sql.DB
}

func InitDB(config *config.Config) (*Database, error) {
	
	url_db := config.DBHost()
	db, err := sql.Open("postgres", url_db)
	if err != nil {
		return nil, fmt.Errorf("sql open: %w", err)
	}
	
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("dp ping: %w", err)
	} 
	return &Database{db}, nil
}




func (db *Database) GetBalanceByUUID(uuid int) (int, error) {
	var balance int
	err := db.db.QueryRow("SELECT balance FROM wallet WHERE wallet_id = $1;", uuid).Scan(&balance)
	return balance, err

}


func (db *Database) CreateWalletByUUID(uuid int, amount int) error {
	_, err:= db.db.Exec("INSERT INTO wallet VALUES ($1, $2);", uuid, amount)
	return err
}


func (db *Database) Deposit(transaction model.Transaction) error {
	balance, err := db.GetBalanceByUUID(transaction.WalletId)
	if err != nil {
		// todo: использовать errors.Is
		if err.Error() == "sql: no rows in result set" {
			db.CreateWalletByUUID(transaction.WalletId, transaction.Amount)
			return nil
 		}
		return err
	}
	_, err = db.db.Exec("UPDATE wallet SET balance = $1 WHERE wallet_id = $2;", balance+transaction.Amount, transaction.WalletId)
	//  todo: использовать update on conflict
	// db.Exec("INSERT into wallet(wallet_id, balance) VALUES ($1, $2)", balance+transaction.Amount, transaction.WalletId)
	return err
}

func (db *Database) Withdraw(transaction model.Transaction) error { 
	balance, err := db.GetBalanceByUUID(transaction.WalletId)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			db.CreateWalletByUUID(transaction.WalletId, transaction.Amount)
			return errors.New("not enough money")
 		}
		return err
	}

	// todo: правила бизнес логики проверять перед функциями database
	if balance < transaction.Amount {
		return errors.New("not enough money")
	} else {

		_, err = db.db.Exec("UPDATE wallet SET balance = $1 WHERE wallet_id = $2;", balance-transaction.Amount, transaction.WalletId)
		return err
	}
}


func (db *Database) Close() error {
	return db.db.Close()
}