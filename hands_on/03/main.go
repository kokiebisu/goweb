package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.html"))
}

func dog(w http.ResponseWriter, req *http.Request) {
	tpl.Execute(w, nil)
}

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	http.HandleFunc("/dog/", dog)
	http.ListenAndServe(":8080", nil)
}