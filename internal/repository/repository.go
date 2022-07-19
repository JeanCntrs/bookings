package repository

import (
	"time"

	"github.com/JeanCntrs/bookings/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(r models.Reservation) (int, error)
	InsertRoomRestriction(rr models.RoomRestriction) error
	SearchAvailabilityByDates(start, end time.Time, roomID int) (bool, error)
}
