package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/index.gohtml"))
}

func index(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("photos")

	if err != nil {
		c = &http.Cookie{
			Name:  "photos",
			Value: "",
		}
	}

	if req.Method == http.MethodPost {
		// get file, header and error
		f, fh, err := req.FormFile("new_post")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer f.Close()

		// get file extension
		ext := strings.Split(fh.Filename, ".")[1]

		h := sha1.New()

		io.Copy(h, f)

		// create new file name
		fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext

		// create new file

		// get current working directory
		cwd, err := os.Getwd()

		if err != nil {
			fmt.Println(err)
		}

		// create file path
		path := filepath.Join(cwd, "public", "images", fname)

		// create new file
		nf, err := os.Create(path)

		if err != nil {
			fmt.Println(err)
		}

		defer nf.Close()

		// reset
		f.Seek(0, 0)

		io.Copy(nf, f)

		if c.Value != "" {
			c.Value += "|"
		}

		c.Value += fname
	}

	xs := strings.Split(c.Value, "|")

	http.SetCookie(w, c)

	err = tpl.ExecuteTemplate(w, "index.gohtml", xs)

	if err != nil {
		fmt.Println(err)
	}

}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
