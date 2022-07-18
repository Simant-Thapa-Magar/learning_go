package main

import (
	"fmt"
	"net/http"
)

func PrintSth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Hello world !")
		next.ServeHTTP(w, req)
	})
}
