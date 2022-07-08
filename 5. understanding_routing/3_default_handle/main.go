package main

import (
	"io"
	"net/http"
)

type window int

func (win window) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Windows")
}

type mac int

func (m mac) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Mactintosh")
}

func main() {
	var w window
	var m mac

	http.Handle("/window", w)
	http.Handle("/mac", m)

	http.ListenAndServe(":8080", nil)
}
