package main

import (
	"io"
	"net/http"
)

type momo int

func (m momo) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Momo")
}

type chowmein int

func (c chowmein) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Chowmein")
}

func main() {
	var m momo
	var c chowmein

	mux := http.NewServeMux()
	mux.Handle("/momo", m)
	mux.Handle("/chowmein", c)

	http.ListenAndServe(":8080", mux)
}
