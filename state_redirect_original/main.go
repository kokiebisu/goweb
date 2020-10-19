package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", handle)
	http.HandleFunc("/dog", dog)
	http.ListenAndServe(":8080", nil)
}

func handle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Location", "/dog")
	// writeheader sends an http response header with the provided status code
	// status code 303 (see other) will change the http method to GET when being redirected
	w.WriteHeader(http.StatusSeeOther)
}

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "successfully redirect to dog")
}