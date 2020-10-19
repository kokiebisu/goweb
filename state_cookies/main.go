package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", addCookies)
	http.HandleFunc("/read", readCookies)
	http.ListenAndServe(":8080", nil)
}

func addCookies(w http.ResponseWriter, req *http.Request) {
	fmt.Println("entered")
	http.SetCookie(w, &http.Cookie{
		Name: "cookie1",
		Value: "value1",
	})
	http.SetCookie(w, &http.Cookie{
		Name: "cookie2",
		Value: "Value2",
	})
	http.SetCookie(w, &http.Cookie{
		Name: "cookie3",
		Value: "value3",
	})
	io.WriteString(w, "added cookies")
}

func readCookies(w http.ResponseWriter, req *http.Request) {
	c1, err := req.Cookie("cookie1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	c2, err := req.Cookie("cookie2")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	c3, err := req.Cookie("cookie3")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	cookies := []*http.Cookie{c1, c2, c3}
	tpl := template.Must(template.ParseFiles("index.html"))
	tpl.Execute(w, cookies)

}