package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

var fm = template.FuncMap{
	"formatTime": dayMonthYear,
}

func dayMonthYear(t time.Time) string {
	return t.Format("02-01-2006")
}

func init() {
	tpl = template.Must(template.New("index.gohtml").Funcs(fm).ParseFiles("index.gohtml"))
}

func main() {
	err := tpl.Execute(os.Stdout, time.Now())

	if err != nil {
		log.Fatalln(err)
	}
}
