package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()
	// using infinite loop to accept eternally
	for {
		// when it receives a tcp protocol it accepts it and creates a connection
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		// just prints to the command line
		fmt.Println("it is connected", conn)
		fmt.Fprintf(conn, "the perfect results")
		// closes the connection
		conn.Close()
	}

}