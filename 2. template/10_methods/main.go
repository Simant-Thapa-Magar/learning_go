package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type person struct {
	Name string
	Age  int
}

func (p person) CanVote() bool {
	if p.Age >= 18 {
		return true
	}
	return false
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	p1 := person{
		"Simant",
		26,
	}

	p2 := person{
		"Rosas",
		8,
	}

	data := []person{p1, p2}

	err := tpl.Execute(os.Stdout, data)

	if err != nil {
		log.Fatalln(err)
	}
}
