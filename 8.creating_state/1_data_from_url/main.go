package main

import (
	"io"
	"net/http"
)

func reader(w http.ResponseWriter, req *http.Request) {
	d := req.FormValue("q")
	io.WriteString(w, "Search query is "+d)
}

func main() {
	http.HandleFunc("/", reader)
	http.ListenAndServe(":8080", nil)
}
