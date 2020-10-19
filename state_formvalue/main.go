package main

import (
	"fmt"
	"net/http"
)

func handle(w http.ResponseWriter, req *http.Request) {
	value := req.FormValue("q")
	fmt.Fprintln(w, "You searched for: ", value)
}

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8080", nil)
}