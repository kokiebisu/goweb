package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
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
	request(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break;
		}
		if i == 0 {
			method := strings.Fields(line)[0]
			fmt.Println(method)
			uri := strings.Fields(line)[1]
			fmt.Println(uri)
			respond(conn, method, uri)
		}
	}
}

func respond(conn net.Conn, method string, uri string) {
	if method == "GET" && uri == "/" {
		body := `<h1>Home</h1>`
		io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
		fmt.Fprint(conn, "Content-Type: text/html\r\n")
		io.WriteString(conn, "\r\n")
		fmt.Fprintf(conn, body)
	}
	if method == "GET" && uri == "/apply" {
		body := `<form method="POST" action="/apply"><input type="submit">Apply</input></form>`
		io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
		fmt.Fprint(conn, "Content-Type: text/html\r\n")
		io.WriteString(conn, "\r\n")
		fmt.Fprintf(conn, body)
	}
	if method == "POST" && uri == "/apply" {
		body := `<h1>apply post</h1>`
		io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
		fmt.Fprint(conn, "Content-Type: text/html\r\n")
		io.WriteString(conn, "\r\n")
		fmt.Fprintf(conn, body)
	}
}