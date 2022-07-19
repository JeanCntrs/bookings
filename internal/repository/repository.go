package repository

import "github.com/JeanCntrs/bookings/internal/models"

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(r models.Reservation) (int, error)
	InsertRoomRestriction(rr models.RoomRestriction) error
}
