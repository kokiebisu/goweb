package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	FirstName string
	LastName string
}

type agent struct {
	person
	License bool
}

func main() {
	tpl, err := template.ParseFiles("go.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	spy := agent{person{FirstName: "Bob", LastName: "Matterson"}, true}

	tpl.ExecuteTemplate(os.Stdout, "go.gohtml", spy)
}