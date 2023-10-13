package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type BookingRoutes struct {
	db *sql.DB
}

func NewBookingRoutes(db *sql.DB) *BookingRoutes {
	return &BookingRoutes{
		db: db,
	}
}

func (br *BookingRoutes) CreateBooking(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON request body into a Booking struct
	var booking Booking
	err := json.NewDecoder(r.Body).Decode(&booking)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the booking data
	// TODO: Implement validation logic

	// Save the booking data to the database
	// TODO: Implement database insertion logic

	// Write a success response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Booking created successfully")
}

func (br *BookingRoutes) GetBookings(w http.ResponseWriter, r *http.Request) {
	// Get user ID from request parameters
	params := mux.Vars(r)
	userID := params["id"]

	// Retrieve bookings for the specified user from the database
	// TODO: Implement database query logic based on user ID

	// Write the booking data to the response as JSON
	bookings := []Booking{
		{ID: "1", UserID: userID, Destination: "Paris", Services: "Hotel"},
		{ID: "2", UserID: userID, Destination: "Tokyo", Services: "Flight"},
	}

	err := json.NewEncoder(w).Encode(bookings)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (br *BookingRoutes) GetBooking(w http.ResponseWriter, r *http.Request) {
	// Get booking ID from request parameters
	params := mux.Vars(r)
	bookingID := params["id"]

	// Retrieve the booking from the database based on the ID
	// TODO: Implement database query logic based on booking ID

	// Write the booking data to the response as JSON
	booking := Booking{
		ID:          bookingID,
		UserID:      "123",     // Example data
		Destination: "Paris",   // Example data
		Services:    "Flight",  // Example data
		PaymentID:   "pay_123", // Example data
	}

	err := json.NewEncoder(w).Encode(booking)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

