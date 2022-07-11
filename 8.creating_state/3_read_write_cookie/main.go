package main

import (
	"fmt"
	"io"
	"net/http"
)

func writeCookie(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my_cookie",
		Value: "Its my cookie",
	})
	io.WriteString(w, "Cookie written. Check developers tool")
}

func readCookie(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my_cookie")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprintln(w, "Your cookie : ", c)

}

func main() {
	http.HandleFunc("/", writeCookie)
	http.HandleFunc("/read", readCookie)
	http.ListenAndServe(":8080", nil)
}
