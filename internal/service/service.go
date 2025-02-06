package service

import (
	"rest-api/internal/database"
	"rest-api/internal/model"
)

// type Service struct {

// }



func/*  (s *Service) */ ProccesTransaction(db *database.Database, transaction model.Transaction) error {
	if err := transaction.Validate(); err != nil {
		return err
	}

	switch transaction.OperationType {
	case model.DEPOSIT:
		if err := db.Deposit(transaction); err != nil {
			return err
		}
	case model.WITHDRAW:
		if err := db.Withdraw(transaction); err != nil {
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
		return 0, err
	}
	
	return balance, nil
}