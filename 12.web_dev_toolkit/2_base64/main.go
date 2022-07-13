package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	text := "Australia is wider than the moon. The moon sits at 3400km in diameter, while Australia's diameter from east to west is almost 4000km."

	encoded := base64.StdEncoding.EncodeToString([]byte(text))

	fmt.Println("Encoded string is ", encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)

	if err != nil {
		log.Fatalln("Its an error")
	}

	fmt.Println("Decoded string is ", string(decoded))

}
