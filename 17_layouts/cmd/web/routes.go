package main

import (
	"lets_try_layouts/pkg/config"
	"lets_try_layouts/pkg/handlers"
	"net/http"

	"github.com/go-chi/chi"
)

func routes(a *config.AppConfig) http.Handler {

	r := chi.NewRouter()
	r.Use(PrintSth)
	r.Use(LoadSession)
	r.Get("/", http.HandlerFunc(handlers.Repo.Home))
	r.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return r
}
