package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(r http.ResponseWriter, rq *http.Request) {
	fmt.Fprintln(r, "Hello world time")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
