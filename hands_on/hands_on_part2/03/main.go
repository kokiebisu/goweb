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
			fmt.Println("line", line)
		}
	}
	// It reaches when you enter a empty string
	fmt.Println("code reached here")
	// this is not following the http protocol so it will not be accepted at the browser
	io.WriteString(conn, "Thanks for reaching out")
}