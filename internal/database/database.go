package database

import (
	"database/sql"
	"errors"
	"fmt"
	"rest-api/internal/config"

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


func (db *Database) UpdateWalletBalance(newBalance int, walletId int) error {
	_, err := db.db.Exec("UPDATE wallet SET balance = $1 WHERE wallet_id = $2;", newBalance, walletId)
	return err
}

func (db *Database) UpdateWalletBalanceOrCreateWallet(balance int, walletId int) error {
	_, err := db.db.Exec(`INSERT INTO wallet (wallet_id, balance)
	VALUES ($1, $2)
	ON CONFLICT (wallet_id) DO UPDATE
	SET balance = wallet.balance + EXCLUDED.balance;`, walletId, balance)
	return err
}





func (db *Database) Close() error {
	return db.db.Close()
}