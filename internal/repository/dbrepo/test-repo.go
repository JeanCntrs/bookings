package dbrepo

import (
	"time"

	"github.com/JeanCntrs/bookings/internal/models"
)

func (pdb *testDBRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts a reservation into the database
func (pdb *testDBRepo) InsertReservation(r models.Reservation) (int, error) {
	return 1, nil
}

// InsertRoomRestriction inserts a room restriction into the database
func (pdb *testDBRepo) InsertRoomRestriction(rr models.RoomRestriction) error {
	return nil
}

// SearchAvailabilityByDatesVyRoomID returns true if availability exists for roomID, and false if no availability
func (pdb *testDBRepo) SearchAvailabilityByDatesVyRoomID(start, end time.Time, roomID int) (bool, error) {
	return false, nil
}

// SearchAvailabilityForAllRooms returns a slice of available rooms, if any, for given date range
func (pdb *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	var rooms []models.Room

	return rooms, nil
}

// GetRoomByID gets a room by id
func (pdb *testDBRepo) GetRoomByID(id int) (models.Room, error) {
	var room models.Room

	return room, nil
}
