package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type BookingHandler struct {
	db *sql.DB
}

type Booking struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Slot      string `json:"slot"`
	CreatedAt string `json:"created_at"`
}

func NewBookingHandler(db *sql.DB) *BookingHandler {
	return &BookingHandler{db}
}

func (bh *BookingHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var booking Booking
	err := json.NewDecoder(r.Body).Decode(&booking)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Failed to read request body: %v", err)
		return
	}

	// TODO: Validate inputs

	stmt, err := bh.db.Prepare("INSERT INTO bookings(id, user_id, slot, created_at) VALUES($1, $2, $3, $4)")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to prepare insert statement: %v", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(booking.ID, booking.UserID, booking.Slot, booking.CreatedAt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to insert booking into database: %v", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Booking created successfully")
}

func (bh *BookingHandler) GetBooking(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookingID := params["id"]

	// TODO: Validate booking ID

	var booking Booking
	row := bh.db.QueryRow("SELECT id, user_id, slot, created_at FROM bookings WHERE id = $1", bookingID)
	err := row.Scan(&booking.ID, &booking.UserID, &booking.Slot, &booking.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Booking not found")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Failed to retrieve booking: %v", err)
		}
		return
	}

	jsonBytes, err := json.Marshal(booking)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to serialize booking: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

// more methods can be added here
