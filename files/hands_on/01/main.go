package main

import (
	"io"
	"net/http"
	"text/template"
)

func foo (w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "foo ran")
}

func dog (w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("dog.gohtml")
	if err != nil {
		http.Error(w, "Error rendering Tempalte", 404);
	}
	tpl.ExecuteTemplate(w, "dog.gohtml", nil)
}

func dogpic (w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "dog.jpg")
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dogpic/", dogpic)
	http.ListenAndServe(":8080", nil)
}