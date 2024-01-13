package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type storage interface {
	CreateAccount(*Account) error
	UpdateAccount(*Account) error
	GetAccounts() ([]*Account, error)
	DeleteAccount(int) error
	GetAccountByID(int) (*Account, error)
}
type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
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

	return &PostgresStore{
		db: db,
	}, nil

}
//Yes I know I can just use an ORM but I want to write raw SQL for this project
func (s *PostgresStore) CreateAccountTable() error {
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

func (s *PostgresStore) CreateAccount(acc *Account) error {

	query := `INSERT INTO account 
	(FirstName,LastName,Number,balance,created_at) 
	VALUES ($1,$2,$3,$4,$5)`
	response, err := s.db.Query(query,
		acc.FirstName,
		acc.LastName,
		acc.Number,
		acc.Balance, acc.CreatedAt)

	if err != nil {
		return err
	}
	log.Printf("%+v", response)
	return nil

}
func (s *PostgresStore) GetAccounts() ([]*Account, error) {
	query := ` SELECT * FROM account`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	accounts := []*Account{}
	for rows.Next() {
		account := new(Account)
		err := rows.Scan(&account.ID, &account.FirstName, &account.LastName, &account.Number, &account.Balance, &account.CreatedAt)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)

	}
	return accounts, nil
}
func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}
func (s *PostgresStore) DeleteAccount(int) error {
	return nil
}
func (s *PostgresStore) GetAccountByID(accId int) (*Account, error) {
	query:=`SELECT * FROM account WHERE id=$1 `
	resp,err := s.db.Query(query,accId)
	if(err != nil){
          return nil,err
	}
	account:=new(Account) //or &Account{}
	err=resp.Scan(&account.ID, &account.FirstName, &account.LastName, &account.Number, &account.Balance, &account.CreatedAt)
	if(err != nil){
		return nil,err
  }

	return account, nil
}
