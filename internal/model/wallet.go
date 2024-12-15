package model

const (
	DEPOSIT  = "DEPOSIT"
	WITHDRAW = "WITHDRAW"
)

type Transaction struct {
	WalletId      int    `json:"walletId"`
	OperationType string `json:"operationType"`
	Amount        int    `json:"amount"`
}

/*
type Wallet struct {
	walletId int `json:"walletId"`
	balance  int `json:"balance"`
} */