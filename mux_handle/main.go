package main

import (
	"fmt"
	"net/http"
)

type handler struct {}

func (h handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "random response")
}

var handler1 handler
var handler2 handler

func main() {
	// *Mux already has the ServeHTTP method so it is of type Handler
	mux := http.NewServeMux()
	// if there is / at the end of the endpoint, it will accept stuff after
	// /dog/it/works/here -> works
	mux.Handle("/dog/", handler1)
	// there is no / after
	// /cat/doesn't/work
	mux.Handle("/cat", handler2)
	http.ListenAndServe(":8080", mux)
}

// Problem with this is that you have to create a new original mux which you don't have to