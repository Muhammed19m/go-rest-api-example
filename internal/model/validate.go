package model

import (
	"errors"
)


var (
	ErrInvalidId = errors.New("invalid id")
	ErrInvalidAmount = errors.New("invalid amount")
	ErrInvalidOperationType = errors.New("invalid operation type")

)



func (tran *Transaction) Validate() error {
	if tran.WalletId <= 0 {
		return ErrInvalidId
	}
	if tran.Amount < 0 {
		return ErrInvalidAmount
	}
	if tran.OperationType != DEPOSIT && tran.OperationType != WITHDRAW {
		return ErrInvalidOperationType
	}
	return nil
}


func ValidateUUid(uuid int) error {
	if uuid <= 0 {
		return ErrInvalidId
	}
	return nil
}