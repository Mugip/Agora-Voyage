go
// main.go
package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Contact struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var contact Contact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Save contact data securely in the database or send emails

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Thank you for contacting us!"))
}

func main() {
	http.HandleFunc("/api/contact", contactHandler)
	http.ListenAndServe(":8080", nil)
  }
