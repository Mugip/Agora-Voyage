go
package models

import (
	"database/sql"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "username:password@tcp(database-host:port)/database-name")
	if err != nil {
		log.Fatal(err)
	}
}

func CreateUser(username, password string) (*User, error) {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	user := &User{
		Username: username,
		Password: hashedPassword,
	}

	query := "INSERT INTO users (username, password) VALUES (?, ?)"
	result, err := db.Exec(query, user.Username, user.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	user.ID, _ = result.LastInsertId()
	return user, nil
}

func GetUserByID(userID int) (*User, error) {
	user := &User{}

	query := "SELECT * FROM users WHERE id = ?"
	err := db.QueryRow(query, userID).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %v", err)
	}

	return user, nil
}

func Authenticate(username, password string) (*User, error) {
	user := &User{}

	query := "SELECT * FROM users WHERE username = ?"
	err := db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to authenticate user: %v", err)
	}

	err = comparePasswords(user.Password, password)
	if err != nil {
		return nil, fmt.Errorf("authentication failed: %v", err)
	}

	// Clear the password field for security reasons
	user.Password = ""
	return user, nil
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func comparePasswords(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
