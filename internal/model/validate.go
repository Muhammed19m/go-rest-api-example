package model

import "fmt"


var (
	ErrInvalidId = fmt.Errorf("invalid id")
	ErrInvalidAmount = fmt.Errorf("invalid amount")
	ErrInvalidOperationType = fmt.Errorf("invalid operation type")
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