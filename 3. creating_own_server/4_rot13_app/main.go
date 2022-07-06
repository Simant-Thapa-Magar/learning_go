package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	tcp, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Panic(err)
	}

	defer tcp.Close()

	for {
		conn, err := tcp.Accept()

		if err != nil {
			log.Println(err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := strings.ToLower(scanner.Text())
		rt := rot13([]byte(ln))
		fmt.Fprintf(conn, "%s - %s\n", ln, rt)
	}
}

func rot13(bs []byte) []byte {
	var r13 = make([]byte, len(bs))
	for i, v := range bs {
		if v <= 109 {
			r13[i] = v + 13
		} else {
			r13[i] = v - 13
		}
	}
	return r13
}
