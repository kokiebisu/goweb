package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/dog/", dog)
	// strips off the first part
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("assets"))))
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, `<img src="/resources/dog.jpg" />`)
}