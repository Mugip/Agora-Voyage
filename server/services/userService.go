go
package service

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var (
	db         *sql.DB
	statements map[string]*sql.Stmt
)

func init() {
	var err error
	db, err = sql.Open("mysql", "username:password@tcp(database-host:port)/database-name")
	if err != nil {
		log.Fatal(err)
	}

	statements = make(map[string]*sql.Stmt)
	statements["insertUser"], _ = db.Prepare("INSERT INTO users (username, password) VALUES (?, ?)")
	statements["getUserByUsername"], _ = db.Prepare("SELECT * FROM users WHERE username = ?")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := statements["insertUser"].Exec(user.Username, hashedPassword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.ID, _ = result.LastInsertId()

	json.NewEncoder(w).Encode(user)
}

func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dbUser := User{}
	err = statements["getUserByUsername"].QueryRow(user.Username).Scan(&dbUser.ID, &dbUser.Username, &dbUser.Password)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Authentication successful
	// Generate and return a JWT token to the client

	// Your JWT token generation logic goes here

	json.NewEncoder(w).Encode(user)
}
