package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type hotdog int

// handler interface is ServeHTTP(...)
// if you assign this to the type, then it is considered a type Handler (subtype)
func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// parses both queries from url (/?...) GET, and form body (POST)
	err := req.ParseForm()
	// req.PostForm(): this is only for form body
	tpl, err := template.ParseFiles("go.gohtml")
	if err != nil {
		panic(err)
	}
	fmt.Println(req.Form)
	tpl.ExecuteTemplate(res, "go.gohtml",req.Form)
}

func main() {
	var hotdog1 hotdog
	http.ListenAndServe(":8080", hotdog1)
}