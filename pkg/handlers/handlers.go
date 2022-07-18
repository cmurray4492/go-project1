package handlers

import (
	"net/http"

	"github.com/cmurray4492/go-project1/pkg/config"
	"github.com/cmurray4492/go-project1/pkg/models"
	"github.com/cmurray4492/go-project1/pkg/render"
)

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

// NewHandlers sets the repository for the handelrs
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// Perform business logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	// Send data to template

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
