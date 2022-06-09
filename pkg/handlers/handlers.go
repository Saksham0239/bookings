package handlers

import (
	"net/http"

	"github.com/Saksham0239/bookings/pkg/config"
	"github.com/Saksham0239/bookings/pkg/models"
	"github.com/Saksham0239/bookings/pkg/render"
)

//This repo will have all the config and global variables to be used by the handlers
//Repo is the repository used by the handlers
var Repo *Repository

//Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

//NewRepo creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//Route Handlers
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	//putting the ip address to the session
	addr := r.RemoteAddr
	m.App.Session.Put(r.Context(), "ip_address", addr)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)

	stringMap["test"] = "Hello test data"

	//getting the session out and adding to stringMap passing it as a data to out template

	addr := m.App.Session.GetString(r.Context(), "ip_address")

	stringMap["addr"] = addr

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{StringMap: stringMap})
}
