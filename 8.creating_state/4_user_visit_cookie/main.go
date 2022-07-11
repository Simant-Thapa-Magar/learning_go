package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func index(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("visit_count")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "visit_count",
			Value: "0",
		}
	}

	count, err := strconv.Atoi(cookie.Value)

	if err != nil {
		log.Fatalln(err)
	}

	count++
	cookie.Value = strconv.Itoa(count)
	http.SetCookie(w, cookie)
	fmt.Fprintln(w, "You have visited this page ", count, " times")
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
