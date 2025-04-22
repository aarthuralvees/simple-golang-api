package handlers

import (
	"github.com/gorilla/mux"
)

func Handler(r *mux.Router) {
	r.HandleFunc("/user", GetUsers).Methods("GET")
	r.HandleFunc("/user/{id}", GetUser).Methods("GET")
	r.HandleFunc("/user", CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")
}
