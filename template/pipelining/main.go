package main

import (
	"log"
	"os"
	"text/template"
)

func hello(s string) string {
	return "Good Morning " + s
}

var fm = template.FuncMap{
	"gt": hello,
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("hello.gohtml").Funcs(fm).ParseFiles("templates/hello.gohtml"))
}

func main() {
	// if you are using Execute, you have to specify the filename inside template.New()
	err := tpl.Execute(os.Stdout, "henry")
	if err != nil {
		log.Fatalln(err)
	}
}