package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func index(w http.ResponseWriter, req *http.Request) {
	t := getTime()
	tpl, err := template.ParseFiles("index.html")

	if err != nil {
		fmt.Fprintln(w, "Sorry ! something went wrong")
		return
	}

	err = tpl.ExecuteTemplate(w, "index.html", t)

	if err != nil {
		fmt.Fprintln(w, "Sorry ! something went wrong")
	}
}

func updateTime(w http.ResponseWriter, req *http.Request) {
	t := getTime()
	fmt.Fprintln(w, t)
}

func getTime() string {
	return time.Now().Format("02-01-2006 15:04:05")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/last-update", updateTime)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
