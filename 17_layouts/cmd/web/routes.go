package main

import (
	"lets_try_layouts/pkg/config"
	"lets_try_layouts/pkg/handlers"
	"net/http"

	"github.com/bmizerany/pat"
)

func routes(a *config.AppConfig) http.Handler {
	r := pat.New()

	r.Get("/", http.HandlerFunc(handlers.Repo.Home))
	r.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return r
}
