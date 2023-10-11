package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

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
	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/protected", protectedHandler).Methods("GET")
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server started on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))

	}

	
