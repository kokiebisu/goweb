package main

import (
	"log"
	"os"
	"text/template"
)


type course struct {
	Name, Number, Description string
}

type semester struct {
	Term string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}


func main() {
	oneYear := year{
		Fall: semester{Term: "fall", Courses: []course{{"y", "303", "something"}}},
		Spring: semester{Term: "spring", Courses: []course{{"s", "301", "sadfijio"}}},
	}

	tpl, err := template.ParseFiles("go.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(os.Stdout, "go.gohtml", oneYear)
}