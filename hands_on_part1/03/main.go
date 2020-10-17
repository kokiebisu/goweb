package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.html"))
}

func home(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "dog.html", nil)
}

func me(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "me.html", nil)
}

func main() {
	http.HandleFunc("/", http.HandlerFunc(home))
	http.HandleFunc("/dog/", http.HandlerFunc(dog))
	http.HandleFunc("/me/", http.HandlerFunc(me))
	http.ListenAndServe(":8080", nil)
}