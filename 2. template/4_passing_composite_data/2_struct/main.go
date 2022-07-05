package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type person struct {
	Fname string
	Lname string
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	p1 := person{
		"Simant",
		"Thapa",
	}

	p2 := person{
		"Second",
		"Person",
	}

	err := tpl.Execute(os.Stdout, []person{p1, p2})

	if err != nil {
		log.Fatalln(err)
	}
}
