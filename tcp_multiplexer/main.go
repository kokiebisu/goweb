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
		connection, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handle(connection)
	}	
}

func handle(conn net.Conn) {
	defer conn.Close()
	request(conn);
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn);
	for scanner.Scan() {
		line := scanner.Text()
		if i == 0 {
			multiplex(conn, line)
		} 
		// you will need this because empty line would still return true for scanner.Scan()
		// it only stops when EOF done by Ctrl+D
		if line == "" {
			break
		}
		i++
	}
}

func multiplex(conn net.Conn, info string) {
	method := strings.Fields(info)[0]
	uri := strings.Fields(info)[1]
	if method == "GET" && uri == "/" {
		home(conn)
	}
	if method == "GET" && uri == "/about" {
		about(conn)
	}
	if method == "POST" && uri == "/about" {
		aboutPost(conn)
	}
}

func home(conn net.Conn) {
	body := `<!doctype html><html><head charset="UTF-8"><title>TCP SERVER</title></head><body><h1>Welcome to the home!</h1><a href="/about">about</a></body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func about(conn net.Conn) {
	content := `<h2>About page</h2><form method="POST"><input type="submit">Send About</input></form>`
	body := `<!doctype html><html><head charset="UTF-8"><title>TCP SERVER</title></head><body>` + content + `</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func aboutPost(conn net.Conn) {
	content := `<h2>About POST</h2>`
	body := `<!doctype html><html><head charset="UTF-8"><title>TCP SERVER</title></head><body>` + content + `</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

