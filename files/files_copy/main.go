package main

import (
	"io"
	"net/http"
)

func dogPic(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// the image is not being served from our server
	io.WriteString(w, `<img src="https://hips.hearstapps.com/hmg-prod.s3.amazonaws.com/images/dog-puppy-on-garden-royalty-free-image-1586966191.jpg" />`)
}
func main() {
	http.HandleFunc("/dog/", dogPic)
	http.ListenAndServe(":8080", nil)
}