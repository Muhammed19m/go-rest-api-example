package service

import (
	"database/sql"
	"errors"
	"rest-api/internal/database"
	"rest-api/internal/model"
)

var (
	ErrNoEnoughMoney = errors.New("not enough money")
	ErrWalletNotFound = errors.New("wallet not found")
)


// type Service struct {

// }



func/*  (s *Service) */ ProccesTransaction(db *database.Database, transaction model.Transaction) error {
	if err := transaction.Validate(); err != nil {
		return err
	}

	switch transaction.OperationType {
	case model.DEPOSIT:
		if err := deposit(db, transaction); err != nil {
			return err
		}
	case model.WITHDRAW:
		if err := withdraw(db, transaction); err != nil {
			return err
		}
	}

	return nil
}


func GetBalance(db *database.Database, uuid int) (int, error) {
	if err := model.ValidateUUid(uuid); err != nil {
		return 0, err
	}

	balance, err := db.GetBalanceByUUID(uuid) 
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrWalletNotFound
		} else {
			return 0, err
		}
	}
	
	return balance, nil
}



func deposit(db *database.Database, transaction model.Transaction) error {
	err := db.UpdateWalletBalanceOrCreateWallet(transaction.Amount, transaction.WalletId)
	return err
}




func withdraw(db *database.Database, transaction model.Transaction) error {
	tx, err := db.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}
	var balance int
	err = tx.QueryRow("SELECT balance FROM wallet WHERE wallet_id = $1;", transaction.WalletId).Scan(&balance)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrWalletNotFound
 		}
		return err
	}

	if balance < transaction.Amount {
		return ErrNoEnoughMoney
	} else {
		_, err := tx.Exec("UPDATE wallet SET balance = $1 WHERE wallet_id = $2;", balance-transaction.Amount, transaction.WalletId)
		if err == nil {
			tx.Commit()
		}
		return err
	}
}
