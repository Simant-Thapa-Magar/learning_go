package main

import (
	"io"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "text/html;charset=UTF-8")
	io.WriteString(w, "Hello World<br><a href='/pug'>Pug Image Here</a>")
}

func pugHandler(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("pug.jpg")

	if err != nil {
		http.Error(w, "No file found", 404)
		return
	}

	defer f.Close()

	io.Copy(w, f)

}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/pug", pugHandler)

	http.ListenAndServe(":8080", nil)
}
