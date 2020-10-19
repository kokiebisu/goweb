package main

import (
	"fmt"
	"net/http"
	"strconv"
)


func main() {
	http.HandleFunc("/", increment)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func increment(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("count")
	// returns type of error when the cookie which you looked for doesn't exist
	if err == http.ErrNoCookie {
		c = &http.Cookie{
			Name: "count",
			Value: "0",
		}
	}
	count, _ := strconv.Atoi(c.Value)
	count++
	c.Value = strconv.Itoa(count)
	http.SetCookie(w, c)
	fmt.Fprintln(w, "Your count is ", c.Value)
}