package main

import (
	"io"
	"net/http"
)

func pugHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "text/html;charset=UTF-8")
	io.WriteString(w, "<img src='/assets/pug.jpg'>")
}

func main() {
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/", pugHandler)

	http.ListenAndServe(":8080", nil)
}
