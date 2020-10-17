package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handle(conn)
	}

}

func handle(conn net.Conn) {
	defer conn.Close()
	keyword := request(conn)
	respond(conn, keyword)
}

func request(conn net.Conn) string {
	i := 0
	var keyword string
	scanner := bufio.NewScanner(conn);
	for scanner.Scan() {
		line := scanner.Text()
		if i == 1 {
			keyword = strings.Fields(line)[1]
		} 
		// you will need this because empty line would still return true for scanner.Scan()
		// it only stops when EOF done by Ctrl+D
		if line == "" {
			break
		}
		i++
	}
	return keyword
}

func respond(conn net.Conn, keyword string) {
	body := `<!doctype html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World to ` + keyword + `</strong></body></html>` 
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}