package main

import (
	"errors"
	"fmt"
	"log"
)

func divide(n, x float32) (float32, error) {
	var result float32
	if x == 0 {
		return result, errors.New("Divide by 0 ? Seriously ?")
	}

	result = n / x

	return result, nil
}

func main() {
	n1, n2, n3, n4 := 100, 10, 15, 0

	result1, err := divide(float32(n1), float32(n2))

	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Division result ", result1)
	}

	result2, e := divide(float32(n3), float32(n4))

	if e != nil {
		log.Println(e)
	} else {
		fmt.Println("Division result ", result2)
	}
}
