package main

import (
	"io"
	"net/http"
)

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, `<img src="/dogpic" />`)
}

func dogpic(w http.ResponseWriter, req *http.Request) {
	// this is way easier to serve files than serve content
	http.ServeFile(w, req, "dog.jpg")
}

func main() {
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dogpic", dogpic)
	http.ListenAndServe(":8080", nil)
}