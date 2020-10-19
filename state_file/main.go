package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

var content string
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func handle(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		
		// returns a file with reader interface
		r, h, err := req.FormFile("g")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println("file: ", r, " header: ", h, " err: " , err)

		defer r.Close()
		cb, err := ioutil.ReadAll(r)
	
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return 
		}
		
		content = string(cb)

		// UPLOADS FILES

		// retrieves a pointer to file. You can either read, write from it
		// Create(name string) (*File, error)
		f, err := os.Create(filepath.Join("./uploads/", h.Filename))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		defer f.Close()

		_, err = f.Write(cb)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}


		w.Header().Set("Content-Type", "text/html")
	}
	tpl.Execute(w, content)
}

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8080", nil)
}