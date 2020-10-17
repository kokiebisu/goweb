package main

import (
	"net/http"
	"net/url"
	"text/template"
)

type handler struct {
}

func (h handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		panic(err)
	}
	// helper function that gets the value based on the key
	name := req.FormValue("name")
	data := struct{
		Method string
		URL *url.URL
		Header http.Header
		Body string
		ContentLength int64
		Host string
	}{
		req.Method,
		req.URL,
		req.Header,
		name,
		req.ContentLength,
		req.Host,
	}
	tpl.ExecuteTemplate(w, "index.html", data)
}

var server handler
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	http.ListenAndServe(":8080", server)
}