package handlers

import (
	"net/http"

	"github.com/bjorndonald/bookings/pkg/config"
	"github.com/bjorndonald/bookings/pkg/models"
	"github.com/bjorndonald/bookings/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.go.tmpl", &models.TemplateData{})
	// fmt.Fprintf(w, "This is the home page")
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP

	// m.App.Session
	render.RenderTemplate(w, "about.go.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
	// sum := addValues(2, 2)
	// _, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page 2 + 2 is %d", sum))
}
