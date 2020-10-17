package main

import (
	"fmt"
	"net/http"
)

type handler struct {}

func (h handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("KenHeader", "This is something I made up")
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "some random text here")
}

// initializes to zero value. Here it is {}
var h handler

func main() {
	fmt.Println(h)
	http.ListenAndServe(":8080", h)
}