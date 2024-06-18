package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddress string
}

func NewAPIServer(listenAddress string) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandlefunc(s.handleAccount))

	router.HandleFunc("/account/{id}", makeHTTPHandlefunc(s.handleGetAccount))

	log.Println("Starting server on", s.listenAddress)

	http.ListenAndServe(s.listenAddress, router)
}

func (a *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return a.handleGetAccount(w, r)
	}
	if r.Method == "POST" {
		return a.handleCreateAccount(w, r)
	}
	if r.Method == "DELETE" {
		return a.handleDeleteAccount(w, r)
	}

	return fmt.Errorf("unsupported method %s", r.Method)
}

func (a *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]

	fmt.Println("id", id)

	return WriteJSON(w, http.StatusOK, &Account{})
}

func (a *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (a *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (a *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func makeHTTPHandlefunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}
