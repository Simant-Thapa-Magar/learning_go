package handlers

import (
	"lets_try_layouts/pkg/config"
	"lets_try_layouts/pkg/model"
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
	r.App.Session.Put(req.Context(), "remote_ip", req.RemoteAddr)
	renderer.Render(w, "home.page.tml", &model.TemplateData{})
}

func (r *Repository) About(w http.ResponseWriter, req *http.Request) {
	remoteIp := r.App.Session.GetString(req.Context(), "remote_ip")
	renderer.Render(w, "about.page.tml", &model.TemplateData{RemoteAddress: remoteIp})
}
