package handlers

import (
	"net/http"

	"github.com/JeanCntrs/bookings/pkg/config"
	"github.com/JeanCntrs/bookings/pkg/models"
	"github.com/JeanCntrs/bookings/pkg/render"
)

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// Repo the repository used by the handlers
var Repo *Repository

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for new handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (rp *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	rp.App.Session.Put(r.Context(), "remoteIp", remoteIp)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (rp *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIp := rp.App.Session.GetString(r.Context(), "remoteIp")
	stringMap["remoteIp"] = remoteIp

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{StringMap: stringMap})
}

// Reservation renders the make a reservation page and display form
func (rp *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "make-reservation.page.html", &models.TemplateData{})
}

// Generals renders the room page
func (rp *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "generals.page.html", &models.TemplateData{})
}

// Majors renders the room page
func (rp *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "majors.page.html", &models.TemplateData{})
}

// Availability renders the search availability page
func (rp *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "search-availability.page.html", &models.TemplateData{})
}

// Contact renders the contact page
func (rp *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.html", &models.TemplateData{})
}
