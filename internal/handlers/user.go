package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aarthuralvees/simple-go-api/store"
	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := store.AllUsers()
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, "Something went wrong", 500)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	id, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, "Invalid id", 400)
	}

	user, err := store.FindUser(id)
	if err != nil {
		http.Error(w, "Something went wrong", 404)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, "Something went wrong", 500)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user store.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Bad request", 500)
		return
	}

	user, err = store.NewUser(user)
	if err != nil {
		http.Error(w, "Unable to create user", 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	id, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, "Invalid id", 400)
	}

	err = store.KillUser(id)

	if err != nil {
		http.Error(w, "unable to delete this user", 400)
	}
}
