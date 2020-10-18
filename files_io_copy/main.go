package main

import (
	"io"
	"net/http"
	"os"
)

func dogpic(w http.ResponseWriter, req *http.Request) {
	// returns a pointer to the file *File
	// You can read from *File or write
	f, err := os.Open("./golden-retriever-royalty-free-image-506756303-1560962726.jpg")
	if err != nil {
		panic(err)
	}
	io.Copy(w, f)
}

func main() {
	http.HandleFunc("/dog/", dogpic)
	http.ListenAndServe(":8080", nil)
}