package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	tcp, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Panic(err)
	}

	defer tcp.Close()

	for {
		conn, e := tcp.Accept()

		if e != nil {
			log.Println(e)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))

	if err != nil {
		fmt.Println("Connection time out !! Bye")
	}

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say: %s\n", ln)
	}
	defer conn.Close()

	// prints on connection close
	fmt.Println("Code at the end")
}
