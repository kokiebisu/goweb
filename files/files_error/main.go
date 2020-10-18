package main

import (
	"log"
	"net/http"
)

func giveError(w http.ResponseWriter, req *http.Request) {
	// You can give a error response back
	// good thing is you can add the status code with the content
	http.Error(w, "I will purposely give an error", 404)
}

func main() {
	http.HandleFunc("/", giveError)
	log.Fatal(http.ListenAndServe(":8080", nil))
}