package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "home")
}

func dog(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "dog")
}

func me(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Ken")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}