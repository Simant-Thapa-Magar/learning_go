package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type score struct {
	Score1 int
	Score2 int
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	d1 := score{
		10,
		5,
	}

	d2 := score{
		12,
		18,
	}

	d3 := score{
		7,
		7,
	}

	err := tpl.Execute(os.Stdout, []score{d1, d2, d3})

	if err != nil {
		log.Fatalln(err)
	}

}
