go
package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	db *sql.DB
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{db}
}

func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Failed to read request body: %v", err)
		return
	}

	// TODO: Validate inputs

	stmt, err := uh.db.Prepare("INSERT INTO users(id, username, email) VALUES($1, $2, $3)")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to prepare insert statement: %v", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.ID, user.Username, user.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to insert user into database: %v", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User created successfully")
}

func (uh *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["id"]

	// TODO: Validate user ID

	var user User
	row := uh.db.QueryRow("SELECT id, username, email FROM users WHERE id = $1", userID)
	err := row.Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "User not found")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Failed to retrieve user: %v", err)
		}
		return
	}

	jsonBytes, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to serialize user: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

// more methods can be added here
