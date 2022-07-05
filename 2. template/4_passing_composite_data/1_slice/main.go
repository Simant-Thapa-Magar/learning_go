package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	str := []string{"I", "am", "learning", "golang"}
	err := tpl.Execute(os.Stdout, str)

	if err != nil {
		log.Fatal(err)
	}
}
