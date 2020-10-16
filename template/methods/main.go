package main

import (
	"log"
	"os"
	"text/template"
)

// function names must be capital
func (p person) Double() int {
	return p.Age * 2
}

// can be used when it is piped
func (p person) WithArgs(n int) int {
	return n * n
}

// the attributes must be capital
type person struct {
	FirstName, LastName string
	Age int
}

func main() {

	data := person{
		FirstName: "John",
		LastName: "Bush",
		Age: 56,
	}

	tpl, err := template.ParseFiles("go.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(os.Stdout, "go.gohtml", data)
}