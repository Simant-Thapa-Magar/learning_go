package main

import (
	"lets_try_layouts/pkg/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/home", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
