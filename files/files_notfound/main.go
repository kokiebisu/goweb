package main

import "net/http"

func main() {

	// HandlerFunc: Makes a function into handler if it has the signature
	// func(w http.ResponseWriter, req http.Request)
	// NotFoundHandler: return HandlerFunc(NotFound)
	http.Handle("/", http.NotFoundHandler())
}