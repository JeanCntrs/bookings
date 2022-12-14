package dbrepo

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/JeanCntrs/bookings/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (pdb *postgresDBRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts a reservation into the database
func (pdb *postgresDBRepo) InsertReservation(r models.Reservation) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var newID int
	stmt := `insert into reservations (first_name, last_name, email, phone, start_date,
			end_date, room_id, created_at, updated_at)
			values($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id`

	err := pdb.DB.QueryRowContext(ctx, stmt,
		r.FirstName,
		r.LastName,
		r.Email,
		r.Phone,
		r.StartDate,
		r.EndDate,
		r.RoomID,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

// InsertRoomRestriction inserts a room restriction into the database
func (pdb *postgresDBRepo) InsertRoomRestriction(rr models.RoomRestriction) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `insert into room_restrictions (start_date, end_date, room_id, reservation_id,
			created_at, updated_at, restriction_id)
			values($1, $2, $3, $4, $5, $6, $7)`

	_, err := pdb.DB.ExecContext(ctx, stmt,
		rr.StartDate,
		rr.EndDate,
		rr.RoomID,
		rr.ReservationID,
		time.Now(),
		time.Now(),
		rr.RestrictionID,
	)

	if err != nil {
		return err
	}

	return nil
}

// SearchAvailabilityByDatesVyRoomID returns true if availability exists for roomID, and false if no availability
func (pdb *postgresDBRepo) SearchAvailabilityByDatesVyRoomID(start, end time.Time, roomID int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select count(id) from room_restrictions where room_id = $1 and $2 < end_date and $3 > start_date`

	var numRows int
	row := pdb.DB.QueryRowContext(ctx, query, roomID, start, end)
	err := row.Scan(&numRows)
	if err != nil {
		return false, err
	}

	if numRows > 0 {
		return false, nil
	}

	return true, nil
}

// SearchAvailabilityForAllRooms returns a slice of available rooms, if any, for given date range
func (pdb *postgresDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select r.id, r.room_name 
			from rooms r 
			where r.id not in 
			(select room_id from room_restrictions rr where $1 < rr.end_date and $2 > rr.start_date)`

	var rooms []models.Room
	rows, err := pdb.DB.QueryContext(ctx, query, start, end)
	if err != nil {
		return rooms, err
	}

	for rows.Next() {
		var room models.Room

		err := rows.Scan(
			&room.ID,
			&room.RoomName,
		)

		if err != nil {
			return rooms, err
		}

		rooms = append(rooms, room)
	}

	if err = rows.Err(); err != nil {
		return rooms, err
	}

	return rooms, nil
}

// GetRoomByID gets a room by id
func (pdb *postgresDBRepo) GetRoomByID(id int) (models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "select id, room_name, created_at, updated_at from rooms where id = $1"

	var room models.Room
	row := pdb.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&room.ID,
		&room.RoomName,
		&room.CreatedAt,
		&room.UpdatedAt,
	)

	if err != nil {
		return room, err
	}

	return room, nil
}

// GetUserByID returns a user by id
func (pdb *postgresDBRepo) GetUserByID(id int) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "select id, first_name, last_name, email, password, access_level, created_at, updated_at from users where id = $1"

	var u models.User
	row := pdb.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.Password,
		&u.AccessLevel,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		return u, err
	}

	return u, nil
}

// UpdateUser updates a user in the database
func (pdb *postgresDBRepo) UpdateUser(u models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "update users set first_name = $1, last_name = $2, email = $3, access_level = $4, updated_at = $5"

	_, err := pdb.DB.ExecContext(ctx, query,
		u.FirstName,
		u.LastName,
		u.Email,
		u.AccessLevel,
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}

// Authenticate authenticates a user
func (pdb *postgresDBRepo) Authenticate(email, testPassword string) (int, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "select id, password from users where email = $1"

	var id int
	var hashedPassword string

	row := pdb.DB.QueryRowContext(ctx, query, email)

	if err := row.Scan(&id, &hashedPassword); err != nil {
		return id, "", err
	}

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(testPassword))

	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, "", errors.New("incorrect password")
	} else if err != nil {
		return 0, "", err
	}

	return id, hashedPassword, nil
}

// AllReservations returns a slice of all reservations
func (pdb *postgresDBRepo) AllReservations() ([]models.Reservation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var reservations []models.Reservation

	query := `
		select r.id, r.first_name, r.last_name, r.email, r.phone, r.start_date, r.end_date, 
		r.room_id, r.created_at, r.updated_at, r.processed, rm.id, rm.room_name 
		from reservations r left join rooms rm on r.room_id = rm.id
		order by r.start_date asc
	`

	rows, err := pdb.DB.QueryContext(ctx, query)
	if err != nil {
		return reservations, err
	}
	defer rows.Close()

	for rows.Next() {
		var r models.Reservation

		err := rows.Scan(
			&r.ID,
			&r.FirstName,
			&r.LastName,
			&r.Email,
			&r.Phone,
			&r.StartDate,
			&r.EndDate,
			&r.RoomID,
			&r.CreatedAt,
			&r.UpdatedAt,
			&r.Processed,
			&r.Room.ID,
			&r.Room.RoomName,
		)

		if err != nil {
			return reservations, err
		}

		reservations = append(reservations, r)
	}

	if err := rows.Err(); err != nil {
		return reservations, err
	}

	return reservations, nil
}

// AllNewReservations returns a slice of all reservations
func (pdb *postgresDBRepo) AllNewReservations() ([]models.Reservation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var reservations []models.Reservation

	query := `
		select r.id, r.first_name, r.last_name, r.email, r.phone, r.start_date, r.end_date, 
		r.room_id, r.created_at, r.updated_at, rm.id, rm.room_name 
		from reservations r left join rooms rm on r.room_id = rm.id
		where processed = 0
		order by r.start_date asc
	`

	rows, err := pdb.DB.QueryContext(ctx, query)
	if err != nil {
		return reservations, err
	}
	defer rows.Close()

	for rows.Next() {
		var r models.Reservation

		err := rows.Scan(
			&r.ID,
			&r.FirstName,
			&r.LastName,
			&r.Email,
			&r.Phone,
			&r.StartDate,
			&r.EndDate,
			&r.RoomID,
			&r.CreatedAt,
			&r.UpdatedAt,
			&r.Room.ID,
			&r.Room.RoomName,
		)

		if err != nil {
			return reservations, err
		}

		reservations = append(reservations, r)
	}

	if err := rows.Err(); err != nil {
		return reservations, err
	}

	return reservations, nil
}

// GetReservationByID returns one reservation by ID
func (pdb *postgresDBRepo) GetReservationByID(id int) (models.Reservation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var res models.Reservation

	query := `
		select r.id, r.first_name, r.last_name, r.email, r.phone, r.start_date, r.end_date,
		r.room_id, r.created_at, r.updated_at, r.processed, rm.id, rm.room_name
		from reservations r
		left join rooms rm on r.room_id = rm.id
		where r.id = $1
	`

	row := pdb.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&res.ID,
		&res.FirstName,
		&res.LastName,
		&res.Email,
		&res.Phone,
		&res.StartDate,
		&res.EndDate,
		&res.RoomID,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.Processed,
		&res.Room.ID,
		&res.Room.RoomName,
	)

	if err != nil {
		return res, err
	}

	return res, nil
}

// UpdateReservation updates a reservation in the database
func (pdb *postgresDBRepo) UpdateReservation(r models.Reservation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "update reservations set first_name = $1, last_name = $2, email = $3, phone = $4, updated_at = $5 where id = $6"

	_, err := pdb.DB.ExecContext(ctx, query,
		r.FirstName,
		r.LastName,
		r.Email,
		r.Phone,
		time.Now(),
		r.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

// DeleteReservation deletes one reservation by id
func (pdb *postgresDBRepo) DeleteReservation(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "delete from reservations where id = $1"

	_, err := pdb.DB.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}

	return nil
}

// UpdateProcessedForReservation updates processed for a reservation by id
func (pdb *postgresDBRepo) UpdateProcessedForReservation(id, processed int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "update reservations set processed = $1 where id = $2"

	_, err := pdb.DB.ExecContext(ctx, query, processed, id)

	if err != nil {
		return err
	}

	return nil
}

func (pdb *postgresDBRepo) AllRooms() ([]models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var rooms []models.Room

	query := `select id, room_name, created_at, updated_at from rooms order by room_name`

	rows, err := pdb.DB.QueryContext(ctx, query)
	if err != nil {
		return rooms, err
	}
	defer rows.Close()

	for rows.Next() {
		var rm models.Room
		err := rows.Scan(
			&rm.ID,
			&rm.RoomName,
			&rm.CreatedAt,
			&rm.UpdatedAt,
		)

		if err != nil {
			return rooms, err
		}

		rooms = append(rooms, rm)
	}

	if err := rows.Err(); err != nil {
		return rooms, err
	}

	return rooms, nil
}

// GetRestrictionsForRoomByDate returns restrictions for a room by date range
func (pdb *postgresDBRepo) GetRestrictionsForRoomByDate(roomID int, start, end time.Time) ([]models.RoomRestriction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var restrictions []models.RoomRestriction

	query := `
		select id, coalesce(reservation_id, 0), restriction_id, room_id, start_date, end_date
		from room_restrictions where $1 < end_date and $2 >= start_date and room_id = $3
	`

	rows, err := pdb.DB.QueryContext(ctx, query, start, end, roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var rr models.RoomRestriction
		err := rows.Scan(
			&rr.ID,
			&rr.ReservationID,
			&rr.RestrictionID,
			&rr.RoomID,
			&rr.StartDate,
			&rr.EndDate,
		)

		if err != nil {
			return nil, err
		}

		restrictions = append(restrictions, rr)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return restrictions, nil
}

// InsertBlockForRoom inserts a room restriction
func (pdb *postgresDBRepo) InsertBlockForRoom(id int, startDate time.Time) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `insert into room_restrictions (start_date, end_date, room_id, restriction_id, created_at, updated_at) 
		values ($1, $2, $3, $4, $5, $6)`

	_, err := pdb.DB.ExecContext(ctx, query, startDate, startDate.AddDate(0, 0, 1), id, 2, time.Now(), time.Now())
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// DeleteBlockByID deletes a room restriction
func (pdb *postgresDBRepo) DeleteBlockByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `delete from room_restrictions where id = $1`

	_, err := pdb.DB.ExecContext(ctx, query, id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
