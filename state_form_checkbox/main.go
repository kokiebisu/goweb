package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

type data struct {
	Content string
	Checkbox bool
}

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8080", nil)
}

func handle(w http.ResponseWriter, req *http.Request) {
	content := req.FormValue("content")
	checkbox := req.FormValue("checkbox") == "on"
	retrieved := data{content, checkbox}
	tpl.Execute(w, retrieved)
}