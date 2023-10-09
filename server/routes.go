go
package main

import (
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the home page!"))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// Implement your login logic here
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	// Implement your protected route logic here
