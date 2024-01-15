package main

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
}
type TransferRequest struct {
	ToAccount int   `json:"toAccount"`
	Amount    int64 `json:"amount"`
}

// add email
type LoginRequest struct {
	Number   int64  `json:"number"`
	Password string `json:"password"`
}

type Account struct {
	ID                int       `json:"id"`
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	Balance           int64     `json:"balance"`
	Number            int64     `json:"number"`
	EncryptedPassword string    `json:"_"`
	CreatedAt         time.Time `json:"createdAt"`
}

func newAccount(firstName, LastName, password string) (*Account, error) {
	encPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}
	return &Account{
		FirstName:         firstName,
		LastName:          LastName,
		Number:            int64(rand.Intn(10000000)),
		EncryptedPassword: string(encPass),
		CreatedAt:         time.Now().UTC(),
		Balance:           0,
	}, nil
}
