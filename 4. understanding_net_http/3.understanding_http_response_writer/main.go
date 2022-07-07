package main

import (
	"fmt"
	"net/http"
)

type momo int

func (m momo) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Simant-Header", "Simant's header")
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	fmt.Fprintln(w, "<h1>Its time for Hello World !</h1>")
}

func main() {
	var s momo
	http.ListenAndServe(":8080", s)
}
