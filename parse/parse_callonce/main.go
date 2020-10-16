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
	err := tpl.ExecuteTemplate(os.Stdout, "hello.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}