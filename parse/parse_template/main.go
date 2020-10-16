package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// tpl is like a container which holds all the template which is being parsed
	tpl, err := template.ParseFiles("tpl.gohtml") // multiple arguments (files) allowed
	if err != nil {
		log.Fatalln(err)
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Println("error creating file")
	}

	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
}