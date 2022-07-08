package main

import (
	"io"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(w, "Bhow Bhow")
		break
	case "/cat":
		io.WriteString(w, "Meow Meow")
		break
	default:
		io.WriteString(w, "What should I say?")
	}
}

func main() {
	var s hotdog
	http.ListenAndServe(":8080", s)
}
