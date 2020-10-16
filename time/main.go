package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now() // this returns type Time
	fmt.Println(t)
	tf := t.Format("Mon Jan 2 15:04:05 MST 2006") // this returns type string
	fmt.Println(tf)
}	