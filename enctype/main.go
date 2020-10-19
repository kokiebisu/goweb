package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

func handle(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		// if the form has the enctype of multipart/form-data then it will not be encoded
		// when it reaches the server. therefore it will include content length
		fmt.Println("contentLength", req.ContentLength)
		l := make([]byte, req.ContentLength)
		// returns the number of bytes which was read
		_, err := req.Body.Read(l)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		tpl, err := template.ParseFiles("index.html")
		if err != nil {
			log.Fatal(err)
		}
		tpl.Execute(w, string(l))
	}
	fmt.Println("here")
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, `<form method="post" enctype="multipart/form-data"><input type="file"><input type="submit"></form>`)
}

func main() {
	http.HandleFunc("/", handle)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}