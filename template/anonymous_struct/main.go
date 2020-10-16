package main

import (
	"log"
)

func main() {
	person := struct{
		firstName string
		lastName string
	}{
		"Ken",
		"Okiebisu",
	}
	log.Println(person)
}