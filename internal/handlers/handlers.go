package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/AshRoy29/booking/internal/config"
	"github.com/AshRoy29/booking/internal/models"
	"github.com/AshRoy29/booking/internal/render"
	"log"
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

	render.RenderTemplates(w, r, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "OLA"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplates(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Suites(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, r, "suites.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Double(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, r, "double.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Book(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, r, "book.page.tmpl", &models.TemplateData{})
}

func (m *Repository) PostBook(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("Arrival date is %s and departure date is %s", start, end)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (m *Repository) BookJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available",
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, r, "contact.page.tmpl", &models.TemplateData{})
}
