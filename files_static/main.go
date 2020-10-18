package main

import "net/http"

func main() {
	// if you have a index.html on the same level, it will not show the file system
	// but instead render html by default
	http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
}