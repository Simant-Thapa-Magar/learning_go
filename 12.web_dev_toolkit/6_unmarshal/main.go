package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type preview struct {
	Url    string
	Width  int
	Height int
}

type product struct {
	Id        string
	Type      string
	Name      string
	Image     preview
	Thumbnail preview
}

func index(w http.ResponseWriter, req *http.Request) {
	var p1 product
	js := `{"id":"0001","type":"donut","name":"Cake","image":{"url":"images/0001.jpg","width":200,"height":200},"thumbnail":{"url":"images/thumbnails/0001.jpg","width":32,"height":32}}`

	err := json.Unmarshal([]byte(js), &p1)

	if err != nil {
		fmt.Fprintln(w, "Something went wrong")
	}

	fmt.Fprintln(w, "Check console")

	fmt.Println("Product", p1)
	fmt.Println("type", p1.Type)
	fmt.Println("Image Url ", p1.Image.Url)
	fmt.Println("Thumbnail Height", p1.Thumbnail.Height)
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
