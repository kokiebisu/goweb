package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			panic(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	respond(conn)
}

func respond(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break;
		} else {
			// the following follows the http protocol so the browser will read it
			body := `This if from html`
			io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
			fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
			fmt.Fprint(conn, "Content-Type: text/plain\r\n")
			io.WriteString(conn, "\r\n")
			fmt.Fprintf(conn, body)
		}
	}
	// It reaches when you enter a empty string
	fmt.Println("code reached here")
	io.WriteString(conn, "Thanks for reaching out")
}