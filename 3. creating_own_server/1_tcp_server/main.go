package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	tcp, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Panic(err)
	}

	defer tcp.Close()

	for {
		con, err := tcp.Accept()

		if err != nil {
			log.Println(err)
		}

		io.WriteString(con, "\n Hello from TCP Server \n")
		fmt.Fprintln(con, "How is your day?")
		fmt.Fprintln(con, "Good I hope")

		con.Close()
	}
}
