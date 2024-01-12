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
	//check if conn is still alive
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &postgresStore{
		db: db,
	}, nil

}

func (s *postgresStore) createAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS account (
		id serial primary key,
		LastName varchar(255),
		FirstName varchar(255),
		Number serial,
		balance bigint,
		created_at timestamp
	); `
	_, err := s.db.Exec(query)

	return err
}

func (s *postgresStore) createAccount(*Account) error {

	query := `INSERT INTO account (id,FirstName,LastName,Number,balance,created_at) VALUES VALUES($1,$2,$3,$4,$5,$6);`
	_, err := s.db.Exec(query)

	return err
}
func (s *postgresStore) updateAccount(*Account) error {
	return nil
}
func (s *postgresStore) deleteAccount(int) error {
	return nil
}
func (s *postgresStore) getAccountByID(int) (*Account, error) {
	return &Account{}, nil
}
