package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type APIserver struct {
	Addr  string
	store storage
}
type APIError struct {
	Error string
}

// default func signature
type APIFunc func(http.ResponseWriter, *http.Request) error

// modifies to correct function signature needed by HandleFunc
func makeHTTPHandleFunc(f APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			//handle the error
			writeJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}

// this function writes JSON
func writeJSON(w http.ResponseWriter, status int, value any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(value)

}

func newAPIServer(Addr string, store storage) *APIserver {
	return &APIserver{
		Addr:  Addr,
		store: store,
	}
}
func (s *APIserver) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleGetAccountByID))
	log.Printf("Server is listening on port %v", s.Addr)
	err := http.ListenAndServe(s.Addr, router)
	if err != nil {
		log.Fatal("Unable to spin up the server")
	}

}

func (s *APIserver) handleAccount(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetAccount(w, r)
	case "POST":
		return s.handleCreateAccount(w, r)
	case "DELETE":
		return s.handleDeleteAccount(w, r)
	default:
		return fmt.Errorf("method is not supported : %s", r.Method)

	}

}
func (s *APIserver) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	accounts, err := s.store.GetAccounts()
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, accounts)
}

func (s *APIserver) handleGetAccountByID(w http.ResponseWriter, r *http.Request) error {
	//returns a map
	id := mux.Vars(r)["id"]
	//refactor this
	intVar, err := strconv.Atoi(id)
	if(err !=nil){
		return  err//errors.New("the id value entered is not a valid number")
	}
	account,err:=s.store.GetAccountByID(intVar)
	if(err !=nil){
		return err
	}
	return writeJSON(w, http.StatusOK, account)
}

func (s *APIserver) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	//read request body and store it in params
	request := createAccountRequest{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		return err
	}

	if request.FirstName == "" || request.LastName == "" {
		return errors.New("fields cannot be empty")
	}
	//newAccount returns a reference to an account struct that is then passed to createAccount()
	account := newAccount(request.FirstName, request.LastName)
	if err := s.store.CreateAccount(account); err != nil {
		return err
	}

	return writeJSON(w, http.StatusCreated, account)
}

func (s *APIserver) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIserver) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
