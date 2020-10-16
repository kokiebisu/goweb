package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()
	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(connection)
	}
}

func handle(conn net.Conn) {
	// times out after 10 seconds
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		// this will never get called
		log.Println("Connection timeout")
	}
	// scanner is *Scanner
	// conn works as an argument because NewScanner accepts a type Reader. 
	// net.Conn has the interface of Read() so it is of type Reader.
	fmt.Println("new routine")
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text() // ln will be a string
		fmt.Println(ln)
		// sends back the response to whereever it got the message from
		fmt.Fprintf(conn, "This is my reply to you! Welcome to TCP!")
	}
	// At the moment, it never reaches the part below
	defer conn.Close()

	fmt.Println("Seems like you are not listening any more so I will turn off")
}