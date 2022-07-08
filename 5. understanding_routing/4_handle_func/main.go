package main

import (
	"io"
	"net/http"
)

func day(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "What day is today ?")
}

func week(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "There are 7 days in a week")
}

func main() {
	http.HandleFunc("/day", day)
	http.HandleFunc("/week", week)

	http.ListenAndServe(":8080", nil)
}
