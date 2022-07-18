package main

import (
	"lets_try_layouts/pkg/config"
	"lets_try_layouts/pkg/handlers"
	"lets_try_layouts/pkg/renderer"
	"log"
	"net/http"
)

func main() {
	var app config.AppConfig

	tc, err := renderer.CreateTemplateCache()

	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc
	app.UseCache = true

	renderer.NewTemplate(&app)

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	http.HandleFunc("/home", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
