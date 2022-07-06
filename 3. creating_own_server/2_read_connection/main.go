package main

import (
	"bufio"
	"fmt"
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
		conn, e := tcp.Accept()

		if e != nil {
			log.Println(e)
		}

		readConn(conn)
	}
}

func readConn(c net.Conn) {
	s := bufio.NewScanner(c)

	for s.Scan() {
		ln := s.Text()
		fmt.Println(ln)
	}

	defer c.Close()

	// Will never reach below
	fmt.Println("I am not supposed to be print while connection is active")
}
