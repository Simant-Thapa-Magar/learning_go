package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	content, err := template.ParseFiles("index.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

	nf, err := os.Create("index.html")

	if err != nil {
		log.Println("error creating file ", err)
	}
	defer nf.Close()

	err = content.Execute(nf, nil)

	if err != nil {
		log.Fatalln(err)
	}
}
