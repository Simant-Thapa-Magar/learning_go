package main

import (
	"lets_try_layouts/pkg/config"
	"lets_try_layouts/pkg/handlers"
	"lets_try_layouts/pkg/renderer"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

var session *scs.SessionManager

func main() {
	var app config.AppConfig

	session = scs.New()

	session.Lifetime = 24 * time.Hour
	session.Cookie.HttpOnly = true
	session.Cookie.Secure = false
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Persist = false

	tc, err := renderer.CreateTemplateCache()

	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc
	app.UseCache = true
	app.Session = session

	renderer.NewTemplate(&app)

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: routes(&app),
	}

	_ = srv.ListenAndServe()
}
