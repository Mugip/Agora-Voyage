package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"
)

type BookingController struct {
	db *sql.DB // MySQL database connection
	// Include any other necessary dependencies or services here
}

// POST /bookings
func (bc *BookingController) CreateBooking(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to obtain the booking details
	var newBooking Booking
	err := json.NewDecoder(r.Body).Decode(&newBooking)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Perform validation on the booking details
	if err := bc.validateBooking(newBooking); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Store the booking in the database
	err = bc.storeBooking(newBooking)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the created booking in the response
	jsonResponse(w, newBooking)
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
}

// Utility function to validate booking details
func (bc *BookingController) validateBooking(booking Booking) error {
	// Add your custom validation logic here
	// For example, check if required fields are present and have valid values
	if booking.UserID == "" {
		return errors.New("UserID is required")
	}

	return nil
}

// Utility function to store booking in MySQL database
func (bc *BookingController) storeBooking(booking Booking) error {
	// Prepare the SQL statement
	stmt, err := bc.db.Prepare("INSERT INTO bookings (user_id, start_time, end_time) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement with the booking details
	_, err = stmt.Exec(booking.UserID, booking.StartTime, booking.EndTime)
	if err != nil {
		return err
	}

	return nil
}
