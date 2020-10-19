package main

import (
	"html/template"
	"log"
	"net/http"
)

// Cookies can be used to recognize which user came to the website
// By storing cookies, the user doesn't have to add their personal information onto the URL
// If it is https, then it decreases the stress of having to exploit information
// urls can be dangerous since some else can add credentials to their url
// cookies cannot be added from the client side but only be deleted. That is what brings additional
// security
func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8080", nil)
}


func handle(w http.ResponseWriter, req *http.Request) {
	cookie := &http.Cookie{
		Name: "first-cookie",
		Value: "congrats",
	}
	http.SetCookie(w, cookie)
	tpl := template.Must(template.ParseFiles("index.html"))
	c, err := req.Cookie("first-cookie")
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	tpl.Execute(w, c)
}
