package models

import (
	"database/sql"
	"fmt"
	"time"
)

type Booking struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "username:password@tcp(database-host:port)/database-name")
	if err != nil {
		log.Fatal(err)
	}
}

func CreateBooking(userID int, startTime, endTime time.Time) (*Booking, error) {
	booking := &Booking{
		UserID:    userID,
		StartTime: startTime,
		EndTime:   endTime,
	}

	query := "INSERT INTO bookings (user_id, start_time, end_time) VALUES (?, ?, ?)"
	result, err := db.Exec(query, booking.UserID, booking.StartTime, booking.EndTime)
	if err != nil {
		return nil, fmt.Errorf("failed to create booking: %v", err)
	}

	booking.ID, _ = result.LastInsertId()
	return booking, nil
}

func GetBookingByID(bookingID int) (*Booking, error) {
	booking := &Booking{}

	query := "SELECT * FROM bookings WHERE id = ?"
	err := db.QueryRow(query, bookingID).Scan(&booking.ID, &booking.UserID, &booking.StartTime, &booking.EndTime)
	if err != nil {
		return nil, fmt.Errorf("failed to get booking: %v", err)
	}

	return booking, nil
}
