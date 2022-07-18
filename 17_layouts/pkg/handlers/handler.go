package handlers

import (
	"lets_try_layouts/pkg/renderer"
	"net/http"
)

func Home(w http.ResponseWriter, req *http.Request) {
	renderer.Render(w, "home.page.tml")
}

func About(w http.ResponseWriter, reeq *http.Request) {
	renderer.Render(w, "about.page.tml")
}
