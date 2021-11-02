package handlers

import (
	"github.com/AshRoy29/booking/pkg/config"
	"github.com/AshRoy29/booking/pkg/models"
	"github.com/AshRoy29/booking/pkg/render"
	"net/http"
)

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the type repository
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

	render.RenderTemplates(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "OLA"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplates(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Suites(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "suites.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Double(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "double.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Book(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "book.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "contact.page.tmpl", &models.TemplateData{})
}
