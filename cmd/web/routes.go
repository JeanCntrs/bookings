package main

import (
	"net/http"

	"github.com/JeanCntrs/bookings/internal/config"
	"github.com/JeanCntrs/bookings/internal/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(noSurf)
	mux.Use(sessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/make-reservation", handlers.Repo.Reservation)
	mux.Get("/generals-quarters", handlers.Repo.Generals)
	mux.Get("/majors-suite", handlers.Repo.Majors)
	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/reservation-summary", handlers.Repo.ReservationSummary)
	mux.Get("/choose-room/{id}", handlers.Repo.ChooseRoom)
	mux.Get("/book-room", handlers.Repo.BookRoom)
	mux.Get("/user/login", handlers.Repo.ShowLogin)
	mux.Get("/user/logout", handlers.Repo.Logout)

	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)
	mux.Post("/make-reservation", handlers.Repo.PostReservation)
	mux.Post("/user/login", handlers.Repo.PostShowLogin)

	mux.Route("/admin", func(r chi.Router) {
		r.Use(Auth)
		r.Get("/dashboard", handlers.Repo.AdminDashboard)

		r.Get("/reservations-new", handlers.Repo.AdminNewReservations)
		r.Get("/reservations-all", handlers.Repo.AdminAllReservations)
		r.Get("/reservations-calendar", handlers.Repo.AdminReservationsCalendar)
		r.Get("/reservations/{src}/{id}/show", handlers.Repo.AdminShowReservation)
		r.Get("/process-reservation/{src}/{id}/do", handlers.Repo.AdminProcessReservation)
		r.Get("/delete-reservation/{src}/{id}/do", handlers.Repo.AdminDeleteReservation)

		r.Post("/reservations/{src}/{id}", handlers.Repo.AdminPostShowReservation)
		r.Post("/reservations-calendar", handlers.Repo.AdminPostReservationsCalendar)
	})

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
