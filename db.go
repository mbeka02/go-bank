package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type storage interface {
	createAccount(*Account) error
	updateAccount(*Account) error
	deleteAccount(int) error
	getAccountByID(int) (*Account, error)
}
type postgresStore struct {
	db *sql.DB
}

func newPostgresStore() (*postgresStore, error) {
	godotenv.Load(".env")
	connStr := os.Getenv("connStr")

	if connStr == "" {
		log.Fatal("connStr is not set")
	}
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &postgresStore{
		db: db,
	}, nil

}

func (s *postgresStore) createAccount(*Account) error {
	return nil
}
func (s *postgresStore) updateAccount(*Account) error {
	return nil
}
func (s *postgresStore) deleteAccount(int) error {
	return nil
}
func (s *postgresStore) getAccountByID(int) (*Account,error) {
	return &Account{},nil
}
