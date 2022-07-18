package renderer

import (
	"bytes"
	"html/template"
	"lets_try_layouts/pkg/config"
	"lets_try_layouts/pkg/model"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

func Render(w http.ResponseWriter, t string, d *model.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		var err error
		tc, err = CreateTemplateCache()
		if err != nil {
			log.Fatal(err)
		}
	}
	ft, ok := tc[t]

	if !ok {
		log.Fatal("No template found")
	}

	buff := new(bytes.Buffer)

	_ = ft.Execute(buff, d)

	buff.WriteTo(w)
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	templateCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tml")

	if err != nil {
		return templateCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return templateCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tml")

		if err != nil {
			return templateCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tml")
			if err != nil {
				return templateCache, err
			}
		}

		templateCache[name] = ts
	}
	return templateCache, nil
}
