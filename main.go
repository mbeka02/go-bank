package main

import (
	"log"
)



func main() {

	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.CreateAccountTable(); err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("%v",store)
	server := newAPIServer(":3000", store)

	server.Run()
}
