package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type person struct {
	Name   string
	Age    int
	Movies []string
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "You are at foo")
}

func marshalHandler(w http.ResponseWriter, req *http.Request) {
	p1 := person{
		"Chris Hemsworth",
		40,
		[]string{"Thor", "Thor: Dark World", "Thor: Ragnarok", "Thor: Love and Thunder"},
	}

	json, err := json.Marshal(p1)

	if err != nil {
		fmt.Fprintln(w, "Something went wrong")
	}

	w.Header().Set("Content-type", "application/json")
	w.Write(json)
}

func encoderHandler(w http.ResponseWriter, req *http.Request) {
	p2 := person{
		"Chris Evans",
		42,
		[]string{"Captain America: The First Avenger", "Captain America: Winter Soldier", "Captain America: Civil War"},
	}

	w.Header().Set("Content-type", "application/json")

	err := json.NewEncoder(w).Encode(p2)

	if err != nil {
		fmt.Println(w, "Something went wrong")
	}
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/marshal", marshalHandler)
	http.HandleFunc("/encoder", encoderHandler)
	http.ListenAndServe(":8080", nil)
}
