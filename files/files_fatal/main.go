package main

import (
	"log"
	"net/http"
)

func main() {
	// You should wrap with log.Fatal because listenAndServe may return error
	// log.Fatal will log the err and exit with 1
	log.Fatal(http.ListenAndServe(":8080", nil))
}