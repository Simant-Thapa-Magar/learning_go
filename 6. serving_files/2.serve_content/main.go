package main

import (
	"io"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "text/html;charset=utf-8")
	io.WriteString(w, "<h3>Welcome<h3><a href='/pug'>Pug image available here</a>")
}

func pugHandler(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("pug.jpg")

	if err != nil {
		http.Error(w, "File not found !", 404)
	}

	fi, err := f.Stat()

	if err != nil {
		http.Error(w, "File not found", 404)
	}

	http.ServeContent(w, req, fi.Name(), fi.ModTime(), f)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/pug", pugHandler)

	http.ListenAndServe(":8080", nil)
}
