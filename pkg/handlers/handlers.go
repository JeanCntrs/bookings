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
