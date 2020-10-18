package main

import (
	"io"
	"net/http"
	"os"
)

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, `<img src="/dogpic" />`)
}

func dogpic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("dog.jpg")
	if err != nil {
		panic(err)
	}
	stat, err := f.Stat()
	if err != nil {
		panic(err)
	}
	// has more control over how the file will be served but can be cumbersome
	http.ServeContent(w, req,stat.Name(), stat.ModTime(), f)
}

func main() {
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dogpic", dogpic)
	http.ListenAndServe(":8080", nil)
}