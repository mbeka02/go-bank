package main

import "math/rand"

type Account struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Balance   int64  `json:"balance"`
	Number    int64  `json:"number"`
}

func newAccount(firstName, LastName string) *Account {
	return &Account{
		ID:        rand.Intn(10000),
		FirstName: firstName,
		LastName:  LastName,
		Number:    int64(rand.Intn(10000000)),
	}
}
