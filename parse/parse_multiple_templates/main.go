package main

import (
	"log"
	"text/template"
	"os"
)


func main() {
	original, err := template.ParseFiles("yo.gogo")
	if err != nil {
		log.Fatalln(err)
	}
	tpl, err := original.ParseFiles("yo2.gogogo") // you can add additional files to the pointer
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "yo2.gogo", nil) // ExecuteTemplate allow you to specify which template you want to execute

	err = tpl.Execute(os.Stdout, nil) // executes whatever is lined up first

}