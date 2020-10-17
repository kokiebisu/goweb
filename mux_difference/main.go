package main

import (
	"fmt"
	"net/http"
)

func d(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "successful")
}

func main() {
	// this must accept handler
	// HandlerFunc is an adapter (means it can convert any type into a type with a type Handler)
	// adapters are used to add specific interfaces to the type
	// type handler has ServeHTTP method
	http.Handle("/dog/", http.HandlerFunc(d))
	http.ListenAndServe(":8080", nil)
}