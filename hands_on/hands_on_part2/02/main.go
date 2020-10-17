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
		fmt.Println("line", line)
	}
	// It actually never reaches here
	fmt.Println("code reached here")
	io.WriteString(conn, "Thanks for reaching out")
}