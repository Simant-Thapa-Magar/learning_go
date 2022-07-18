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

	srv := &http.Server{
		Addr:    ":8080",
		Handler: routes(&app),
	}

	_ = srv.ListenAndServe()
}
