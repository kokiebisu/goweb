package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template // global scope

type country struct {
	Name string
	Color string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*")) // checks the errors for you. more concise
}

func main() {
	country1 := country{Name: "Korea", Color: "Blue"}
	country2 := country{Name: "Japan", Color: "Red"}
	countries := []country{country1, country2}
	err := tpl.ExecuteTemplate(os.Stdout, "hello.gohtml", countries)
	if err != nil {
		log.Fatalln(err)
	}
}