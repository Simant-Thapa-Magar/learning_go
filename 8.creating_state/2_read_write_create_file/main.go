package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, req *http.Request) {
	var s string

	if req.Method == http.MethodPost {
		f, h, e := req.FormFile("q")

		if e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}

		defer f.Close()

		fmt.Println("file ", f, "\nHeader", h, "\nError", e)

		bs, e := ioutil.ReadAll(f)

		if e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}

		s = string(bs)

		fi, e := os.Create("copy.txt")

		if e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}

		defer fi.Close()

		_, e = fi.Write(bs)
		if e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("content-type", "text/html;charset=UTF-8")
	io.WriteString(w, "<form method='POST' enctype='multipart/form-data'><input type='file' name='q'><input type='submit' value='Go Ahead !'/></form><br>"+s)
}

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
