package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type UserRoutes struct {
	db *sql.DB
}

func NewUserRoutes(db *sql.DB) *UserRoutes {
	return &UserRoutes{
		db: db,
	}
}

func (ur *UserRoutes) GetUser(w http.ResponseWriter, r *http.Request) {
	// Get user ID from request parameters
	params := mux.Vars(r)
	userID := params["id"]

	// Retrieve user from the database based on the ID
	// TODO: Implement this based on your database schema

	// Write user data to the response as JSON
	user := User{
		ID:   userID,
		Name: "John Doe", // Example data
	}

	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (ur *UserRoutes) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Get user ID from request parameters
	params := mux.Vars(r)
	userID := params["id"]

	// Decode the JSON request body into a User struct
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update user data in the database based on the ID
	// TODO: Implement this based on your database schema

	// Write a success response
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User %s updated successfully", userID)
}
