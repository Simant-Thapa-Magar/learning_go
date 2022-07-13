package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func getCode(key string, message string) string {
	c := hmac.New(sha256.New, []byte(key))
	c.Write([]byte(message))
	return fmt.Sprintf("%x", c.Sum(nil))
}

func main() {
	s1 := "abcd"
	k1 := "key1"
	s2 := "xyz"
	k2 := "key2"

	fmt.Println("hash for ", s1, " with key ", k1, " is ", getCode(s1, k1))
	fmt.Println("hash for ", s1, " with key ", k1, " is ", getCode(s1, k2))
	fmt.Println("hash for ", s2, " witj key ", k2, " is ", getCode(s2, k1))
	fmt.Println("hash for ", s2, " witj key ", k2, " is ", getCode(s2, k2))
}
