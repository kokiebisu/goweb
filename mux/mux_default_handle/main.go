package main

import (
	"fmt"
	"net/http"
)

type handler struct {}

func (h handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "yo")
}

func main() {
	var handler1 handler
	http.Handle("/dog", handler1)

	http.ListenAndServe(":8080", nil)
}