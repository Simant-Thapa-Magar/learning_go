package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var tpl *template.Template

var fn = template.FuncMap{
	"getDouble":     getDouble,
	"getSquare":     getSquare,
	"getSquareRoot": getSquareRoot,
}

func getDouble(n int) int {
	return n + n
}

func getSquare(n int) float64 {
	return math.Pow(float64(n), 2)
}

func getSquareRoot(n float64) float64 {
	return math.Sqrt(n)
}

func init() {
	tpl = template.Must(template.New("index.gohtml").Funcs(fn).ParseFiles("index.gohtml"))
}

func main() {
	err := tpl.Execute(os.Stdout, 5)

	if err != nil {
		log.Fatalln(err)
	}
}
