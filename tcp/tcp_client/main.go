package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	// func Dial(network, address string) (Conn, error)
	connection, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln("Connection not established")
	}
	defer connection.Close()

	// Reads
	bs, err := ioutil.ReadAll(connection)
	if err != nil {
		log.Fatalln("error")
	}
	fmt.Println("You got", string(bs))

}