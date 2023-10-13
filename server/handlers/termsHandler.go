package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func termsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// Retrieve the terms and conditions content from a secure source (e.g., database or file)

	termsContent := `These are our terms and conditions. Please read them carefully...`

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(termsContent))
}

func main() {
	http.HandleFunc("/api/terms", termsHandler)
	http.ListenAndServe(":8080", nil)
}
