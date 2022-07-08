package main

import (
	"fmt"
	"net/http"
)

func myHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("URL ", req.URL.Path)
	fmt.Fprintln(w, "Go check in terminal")
}

func main() {
	http.HandleFunc("/", myHandler)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
