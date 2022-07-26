package dbrepo

import (
	"errors"
	"time"

	"github.com/JeanCntrs/bookings/internal/models"
)

func (pdb *testDBRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts a reservation into the database
func (pdb *testDBRepo) InsertReservation(r models.Reservation) (int, error) {
	// if the room id is 2, then fail; otherwise, pass
	if r.RoomID == 2 {
		return 0, errors.New("some error")
	}

	return 1, nil
}

// InsertRoomRestriction inserts a room restriction into the database
func (pdb *testDBRepo) InsertRoomRestriction(rr models.RoomRestriction) error {
	if rr.RoomID == 1000 {
		return errors.New("some error")
	}

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

	if id > 2 {
		return room, errors.New("some error")
	}

	return room, nil
}
