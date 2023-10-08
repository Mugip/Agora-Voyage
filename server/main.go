go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"your-package-name/handlers"
)

func main() {
	db := setupDatabase()
	defer db.Close()

	authHandler := handlers.NewAuthHandler(db)

	router := mux.NewRouter()

	// Public routes
	router.HandleFunc("/api/login", authHandler.Login).Methods("POST")

	// Authenticated routes
	authRouter := router.PathPrefix("/api").Subrouter()
	authRouter.Use(ValidateToken)

	authRouter.HandleFunc("/logout", authHandler.Logout).Methods("POST")

	// more routes for other features will be added here

	log.Println("Server is running on port 8080...")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
