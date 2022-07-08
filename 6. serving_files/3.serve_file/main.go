package main

import (
	"io"
	"net/http"
)

func indexHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "text/html;charset=UTF-8")
	io.WriteString(w, "<h3>Welcome Back</h3><a href='/pug'>Want some more pug ?</a>")
}

func pugHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "pug.jpg")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/pug", pugHandler)

	http.ListenAndServe(":8080", nil)
}
