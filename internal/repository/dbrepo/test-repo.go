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

func (pdb *testDBRepo) GetUserByID(id int) (models.User, error) {
	var u models.User

	return u, nil
}

func (pdb *testDBRepo) UpdateUser(u models.User) error {
	return nil
}

func (pdb *testDBRepo) Authenticate(email, testPassword string) (int, string, error) {
	if email == "me@here.ca" {
		return 1, "", nil
	}

	return 0, "", errors.New("some error")
}

func (pdb *testDBRepo) AllReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation

	return reservations, nil
}

func (pdb *testDBRepo) AllNewReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation

	return reservations, nil
}

func (pdb *testDBRepo) GetReservationByID(id int) (models.Reservation, error) {
	var res models.Reservation

	return res, nil
}

func (pdb *testDBRepo) UpdateReservation(r models.Reservation) error {
	return nil
}

func (pdb *testDBRepo) DeleteReservation(id int) error {
	return nil
}

func (pdb *testDBRepo) UpdateProcessedForReservation(id, processed int) error {
	return nil
}

func (pdb *testDBRepo) AllRooms() ([]models.Room, error) {
	var rooms []models.Room

	return rooms, nil
}

func (pdb *testDBRepo) GetRestrictionsForRoomByDate(roomID int, start, end time.Time) ([]models.RoomRestriction, error) {
	var restrictions []models.RoomRestriction

	return restrictions, nil
}

func (pdb *testDBRepo) InsertBlockForRoom(id int, startDate time.Time) error {
	return nil
}

func (pdb *testDBRepo) DeleteBlockByID(id int) error {
	return nil
}
