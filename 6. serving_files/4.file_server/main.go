package main

import (
	"io"
	"net/http"
)

func pugHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "text/html;charset=UTF-8")
	io.WriteString(w, "<img src='pug.jpg'/>")
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/pug", pugHandler)

	http.ListenAndServe(":8080", nil)
}
