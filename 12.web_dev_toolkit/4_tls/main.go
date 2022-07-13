package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello world")
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
}
