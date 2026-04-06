package main

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
}

type APIError struct {
	Error string
}

// //Function signature for API handlers
type apiFunc func(http.ResponseWriter, *http.Request) error

func WriteJSON(w http.ResponseWriter, status int, v any) error{
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func makeHttpHandlerFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			//handle the error
			WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}

func NewAPIServer(listenAddr string) *APIServer{
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run(){
	router := mux.NewRouter()
	router.HandleFunc("/account", makeHttpHandlerFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", makeHttpHandlerFunc(s.handleGetAccount))
	log.Printf("Starting server on %s", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error{
	if r.Method == "GET"{
		return s.handleGetAccount(w, r)
	}
	if r.Method == "POST"{
		return s.handleCreateAccount(w, r)
	}
	if r.Method == "PUT"{
		return s.handleUpdateAccount(w, r)
	}
	if r.Method == "DELETE"{
		return s.handleDeleteAccount(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error{
	return nil
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error{
	id := mux.Vars(r)["id"]
	log.Printf("Getting account with id %s", id)
	account := NewAccount("John", "Doe")
	return WriteJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleUpdateAccount(w http.ResponseWriter, r *http.Request) error{
	return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error{
	return nil
}

func (s *APIServer) handleTransferAccount(w http.ResponseWriter, r *http.Request) error{
	return nil
}