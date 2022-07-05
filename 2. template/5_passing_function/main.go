package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

type player struct {
	Name        string
	Kit         int64
	Nationality string
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func init() {
	tpl = template.Must(template.New("index.gohtml").Funcs(fm).ParseFiles("index.gohtml"))
}

func main() {
	p1 := player{
		"Cristiano Ronaldo",
		7,
		"Portugal",
	}

	p2 := player{
		"Lionel Messi",
		10,
		"Argentina",
	}

	p3 := player{
		"David De Gea",
		1,
		"Spain",
	}

	err := tpl.Execute(os.Stdout, []player{p1, p2, p3})

	if err != nil {
		log.Fatal(err)
	}
}
