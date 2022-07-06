package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Panic(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()

		if err != nil {
			log.Println(err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	io.WriteString(conn, "\n IN-MEMORY DATABASE \n\n"+"USE: \n"+"SET key value \n"+"GET key\n"+"DELETE key\n\n"+"EXAMPLE:\n"+"SET fav golang\n"+"GET fav\n"+"DELETE fav\n\n\n")

	data := make(map[string]string)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		in := scanner.Text()
		fr := strings.Fields(in)

		switch fr[0] {
		case "GET":
			if len(fr) < 2 {
				fmt.Fprintln(conn, "Missing key !!!")
			} else {
				k := fr[1]
				v, ok := data[k]
				if ok {
					fmt.Fprintf(conn, "%s\n", v)
				} else {
					fmt.Fprintf(conn, "Key %s not found\n", k)
				}
			}
			break
		case "SET":
			if len(fr) < 3 {
				fmt.Fprintln(conn, "Missing arguments !!!!")
			} else {
				k := fr[1]
				v := fr[2]
				data[k] = v
				fmt.Fprintf(conn, "%s saved successfully\n", k)
			}
			break
		case "DELETE":
			if len(fr) < 2 {
				fmt.Fprintln(conn, "Missing key !!!")
			} else {
				k := fr[1]
				v, ok := data[k]
				if ok {
					delete(data, k)
					fmt.Fprintf(conn, "[%s]=%s deleted successfully\n", k, v)
				} else {
					fmt.Fprintf(conn, "Key %s not found\n", k)
				}
			}
			break
		default:
			fmt.Fprintln(conn, "Invalid command")
		}

	}
}
