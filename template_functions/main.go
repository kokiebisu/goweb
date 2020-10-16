package main

import (
	"fmt"
	"os"
	"text/template"
)

var tpl *template.Template

func hello() string {
	return "Hello~"
}

func greeting() string {
	return "voila"
}

var funcMap = template.FuncMap{
	"hello": hello,
	"greeting": greeting,
}

func init() {
	tpl = template.Must(template.New("").Funcs(funcMap).ParseFiles("templates/hello.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "hello.gohtml", nil)
	if err != nil {
		fmt.Println("wrong")
	}
}