package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello world from docker container")
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":80", nil)
}
