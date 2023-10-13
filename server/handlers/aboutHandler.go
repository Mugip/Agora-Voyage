package main

import (
	"html/template"
	"net/http"
)

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/about.html"))
	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/about", aboutHandler)
	http.ListenAndServe(":8080", nil)
}
