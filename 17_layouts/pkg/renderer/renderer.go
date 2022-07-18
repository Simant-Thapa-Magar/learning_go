package renderer

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

func Render(w http.ResponseWriter, t string) {
	templates, err := CreateTemplateCache()

	if err != nil {
		log.Fatal(err)
	}
	ft, ok := templates[t]

	if !ok {
		log.Fatal("No template found")
	}

	buff := new(bytes.Buffer)

	_ = ft.Execute(buff, nil)

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
