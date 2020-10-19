package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", handle)
	http.HandleFunc("/redirect", redirect)
	http.HandleFunc("/dog", dog)
	http.ListenAndServe(":8080", nil)
}

func handle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	tpl := template.Must(template.ParseFiles("index.html"))
	tpl.Execute(w, nil)
}

func redirect(w http.ResponseWriter, req *http.Request) {
	// temporary redirect redirects the same http method
	// can be used when you want to redirect the POST method
	http.Redirect(w, req, "/dog", http.StatusTemporaryRedirect)
}

func dog(w http.ResponseWriter, req *http.Request) {
	fmt.Println("method: ", req.Method)
	io.WriteString(w, "hehe")
}