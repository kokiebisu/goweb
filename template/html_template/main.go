package main

import (
	"html/template"
	"log"
	"os"
)


func main() {
	tpl, err := template.ParseFiles("go.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	data := struct{
		Title, Input string
	}{
		Title: "This is the title",
		Input: "<script>Hello</script>", // text/template will recognize this as a script tag but html/template doesn't
	}
	err = tpl.ExecuteTemplate(os.Stdout, "go.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}