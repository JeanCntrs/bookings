package repository

import (
	"time"

	"github.com/JeanCntrs/bookings/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(r models.Reservation) (int, error)
	InsertRoomRestriction(rr models.RoomRestriction) error
	SearchAvailabilityByDatesVyRoomID(start, end time.Time, roomID int) (bool, error)
	SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error)
	GetRoomByID(id int) (models.Room, error)
}
