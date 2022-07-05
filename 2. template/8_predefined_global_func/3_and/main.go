package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type player struct {
	Name      string
	Nation    string
	FreeAgent bool
	Affodable bool
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	p1 := player{
		"Cristiano Ronaldo",
		"Portugal",
		false,
		true,
	}

	p2 := player{
		"Paul Pogba",
		"France",
		true,
		false,
	}

	p3 := player{
		"Paulo Dybala",
		"Argentina",
		true,
		true,
	}

	p4 := player{
		"Frankie De Jong",
		"Netherlands",
		false,
		false,
	}

	err := tpl.Execute(os.Stdout, []player{p1, p2, p3, p4})

	if err != nil {
		log.Fatalln(err)
	}
}
