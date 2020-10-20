package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", handle)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)
	http.ListenAndServe(":8080", nil)
}

func handle(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, `<a href="/set">set</a>`)
}

func set(w http.ResponseWriter, req *http.Request) {
	cookie := &http.Cookie{
		Name: "random",
		Value: "congrats",
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, req, "/read", http.StatusSeeOther)
}

func read(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("random")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, `cookie=`, c.Value)
	fmt.Fprintln(w, `<a href="/expire">Expire</a>`)
}

func expire(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("random")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	c.MaxAge = -1
	http.SetCookie(w, c)
}