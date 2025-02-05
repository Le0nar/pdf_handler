package ticket

import (
	"time"

	"github.com/google/uuid"
)

type Ticket struct {
	ID            uuid.UUID `json:"id" db:"id" validate:"required"`
	PassengerName string    `json:"passenger_name" db:"passenger_name" validate:"required,min=3,max=100"`
	FlightNumber  string    `json:"flight_number" db:"flight_number" validate:"required"`
	Departure     time.Time `json:"departure" db:"departure" validate:"required"`
	Arrival       time.Time `json:"arrival" db:"arrival" validate:"required"`
	From          string    `json:"from" db:"from" validate:"required"`
	To            string    `json:"to" db:"to" validate:"required"`
	SeatNumber    string    `json:"seat_number" db:"seat_number" validate:"required"`
	Class         string    `json:"class" db:"class" validate:"required,oneof=economy business first"`
	Price         float64   `json:"price" db:"price" validate:"required,gt=0"`
}
