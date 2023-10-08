go
package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ProductController struct {
	// Include any necessary dependencies or services here
}

// GET /products
func (pc *ProductController) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	// Fetch all products from the database or any other storage mechanism
	products := []Product{
		{ID: 1, Name: "Product 1", Price: 10.99},
		{ID: 2, Name: "Product 2", Price: 19.99},
		{ID: 3, Name: "Product 3", Price: 8.99},
	}

	jsonResponse(w, products)
}

// POST /products
func (pc *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to obtain the new product details
	var newProduct Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Add the new product to the database or any other storage mechanism
	// ...

	// Return the created product in the response
	jsonResponse(w, newProduct)
}

// Utility function to send JSON response
func jsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
