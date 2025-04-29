package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/UserLeeIsTired/graceful_shutdown/database"
	"github.com/go-chi/chi/v5"
)

func UserRoutes(r chi.Router, db *database.Database) {

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		GetAllUsers(w, r, db)
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		CreateUser(w, r, db)
	})

	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		GetUserById(w, r, db)
	})
}

func GetAllUsers(w http.ResponseWriter, r *http.Request, db *database.Database) {
	users, err := db.GetAllUsers()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	WriteHeader(w, http.StatusOK, users)
}

func GetUserById(w http.ResponseWriter, r *http.Request, db *database.Database) {
	id := chi.URLParam(r, "id")
	user, err := db.GetUserById(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	WriteHeader(w, http.StatusOK, user)
}

func CreateUser(w http.ResponseWriter, r *http.Request, db *database.Database) {
	var user database.User

	// Decode the request body
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Simulate a long-running operation
	time.Sleep(5 * time.Second)

	id, err := db.CreateUser(user.Name)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	WriteHeader(w, http.StatusCreated, id)
}
