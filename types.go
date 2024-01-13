package main

import (
	"math/rand"
	"time"
)

type createAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
type TransferRequest struct {
	ToAccount int `json:"toAccount"`
	Amount int64 `json:"amount"`
}

type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Balance   int64     `json:"balance"`
	Number    int64     `json:"number"`
	CreatedAt time.Time `json:"createdAt"`
}

func newAccount(firstName, LastName string) *Account {
	return &Account{
		FirstName: firstName,
		LastName:  LastName,
		Number:    int64(rand.Intn(10000000)),
		CreatedAt: time.Now().UTC(),
		Balance:   0,
	}
}
