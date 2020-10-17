package main

import (
	"fmt"
	"net/http"
)

func respond(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "successfuly rendered")
}

func main() {
	// second argument accepts a function with arguments (responsewriter, request)
	http.HandleFunc("/dog/", respond)
	// if it is nil, then it creates defaultMux
	http.ListenAndServe(":8080", nil)
}