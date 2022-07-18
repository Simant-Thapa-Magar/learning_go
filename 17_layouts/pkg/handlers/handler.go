package handlers

import (
	"lets_try_layouts/pkg/config"
	"lets_try_layouts/pkg/renderer"
	"net/http"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (r *Repository) Home(w http.ResponseWriter, req *http.Request) {
	renderer.Render(w, "home.page.tml")
}

func (r *Repository) About(w http.ResponseWriter, reeq *http.Request) {
	renderer.Render(w, "about.page.tml")
}
