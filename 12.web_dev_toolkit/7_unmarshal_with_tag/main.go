package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type star struct {
	Name   string `json:"Name"`
	Sex    string `json:"Gender"`
	Galaxy string `json:"Homeworld"`
	Dob    string `json:"Born"`
	Jedi   string `json:"Jedi"`
}

func index(w http.ResponseWriter, req *http.Request) {
	var s1 star
	data := `{"Name":"Anakin","Gender":"male","Homeworld":"Tatooine","Born":"41.9BBY","Jedi":"yes"}`

	err := json.Unmarshal([]byte(data), &s1)

	if err != nil {
		fmt.Fprintln(w, "Something went wrong")
		return
	}

	fmt.Fprintln(w, "Check console for output")

	fmt.Println("Star Warrior ", s1)
	fmt.Println("Name ", s1.Name)
	fmt.Println("Sex ", s1.Sex)
	fmt.Println("From ", s1.Galaxy)
	fmt.Println("Born ", s1.Dob)
	fmt.Println("Jedi ", s1.Jedi)
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
