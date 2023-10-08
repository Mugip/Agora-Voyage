go
package main

import (
	"database/sql"
	"fmt"
)

// RunMigrations runs all the required database migrations
func RunMigrations(db *sql.DB) error {
	migrations := []func(*sql.DB) error{
		createUsersTable,
		createPostsTable,
		// Add other migration functions here
	}

	for _, migration := range migrations {
		err := migration(db)
		if err != nil {
			return err
		}
	}

	return nil
}

// createUsersTable creates the users table in the database
func createUsersTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(50) UNIQUE NOT NULL,
			password VARCHAR(100) NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		)
	`)

	return err
}

// createPostsTable creates the posts table in the database
func createPostsTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS posts (
			id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL,
			title VARCHAR(100) NOT NULL,
			content TEXT NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
		)
	`)

	return err
}

// Add other migration functions for additional tables here
