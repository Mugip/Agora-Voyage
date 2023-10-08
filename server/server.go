go
package main

import (
	"log"
	"net/http"
)

func main() {
	// Create a new instance of our database
	db, err := NewDatabase()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	// Run the required database migrations
	if err := db.RunMigrations(); err != nil {
		log.Fatal("Failed to run database migrations:", err)
	}

	// Create a new instance of a router
	router := NewRouter(db)

	// Specify the routes
	router.HandleFunc("/users", CreateUserHandler).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", GetUserHandler).Methods(http.MethodGet)
	router.HandleFunc("/posts", CreatePostHandler).Methods(http.MethodPost)
	router.HandleFunc("/posts/{id}", GetPostHandler).Methods(http.MethodGet)

	// Start the server on port 8080
	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
