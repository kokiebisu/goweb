package main

import (
	"io"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8080", nil)
}

func handle(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("count")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name: "count",
			Value: "0",
		}
	}
	count, _ := strconv.Atoi(cookie.Value)
	cookie.Value = strconv.Itoa(count + 1)
	http.SetCookie(w, cookie)
	io.WriteString(w, cookie.Value)
}