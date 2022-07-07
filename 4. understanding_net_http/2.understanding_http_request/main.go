package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var tpl *template.Template

type burger int

func (b burger) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	err := req.ParseForm()

	if err != nil {
		log.Fatal(err)
	}

	data := struct {
		Method string
		Body   url.Values
	}{
		req.Method,
		req.Form,
	}
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var hb burger
	http.ListenAndServe(":8080", hb)
}
