package main

import (
	"flag"
	"fmt"
	"log"
)

func seedAccount(s storage, fname, lname, pw string) *Account {
	acc, err := newAccount(fname, lname, pw)

	if err != nil {
		log.Fatal(err)
	}
	if err := s.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("new account details => %+v", acc)
	return acc
}
func seedAccounts(s storage) {
	seedAccount(s, "Michael", "Mbeka", "myPassword1")
	seedAccount(s, "Anthony", "Mbeka", "myPassword2")

}

func main() {
	seed := flag.Bool("seed", false, "seed the db")
	flag.Parse()

	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.CreateAccountTable(); err != nil {
		log.Fatal(err)
	}
	if *seed {
		//Seed stuff
		fmt.Println("...seeding the database")
		seedAccounts(store)
	}

	server := newAPIServer(":3000", store)

	server.Run()
}
