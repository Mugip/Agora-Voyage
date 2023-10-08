go
package service

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Booking struct {
	ID          int64  `json:"id"`
	UserID      int64  `json:"userId"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var (
	db         *sql.DB
	statements map[string]*sql.Stmt
)

func init() {
	var err error
	db, err = sql.Open("mysql", "username:password@tcp(database-host:port)/database-name")
	if err != nil {
		log.Fatal(err)
	}

	statements = make(map[string]*sql.Stmt)
	statements["insertBooking"], _ = db.Prepare("INSERT INTO bookings (userId, title, description) VALUES (?, ?, ?)")
}

func GetAllBookings(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM bookings")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	bookings := make([]Booking, 0)
	for rows.Next() {
		var booking Booking
		err := rows.Scan(&booking.ID, &booking.UserID, &booking.Title, &booking.Description)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		bookings = append(bookings, booking)
	}

	json.NewEncoder(w).Encode(bookings)
}

func AddBooking(w http.ResponseWriter, r *http.Request) {
	var booking Booking
	err := json.NewDecoder(r.Body).Decode(&booking)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := statements["insertBooking"].Exec(booking.UserID, booking.Title, booking.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	booking.ID, _ = result.LastInsertId()

	json.NewEncoder(w).Encode(booking)
}
