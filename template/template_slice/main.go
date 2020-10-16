package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template // global scope

func init() {
	tpl = template.Must(template.ParseGlob("templates/*")) // checks the errors for you. more concise
}

func main() {
	slices := []string{"one", "two", "three"}
	err := tpl.ExecuteTemplate(os.Stdout, "hello.gohtml", slices)
	if err != nil {
		log.Fatalln(err)
	}
}