package main

import (
	"log"
)

//"net/http"

func main() {

	store, err := newPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.createAccountTable(); err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("%v",store)
	server := newAPIServer(":3000", store)

	server.Run()
}
