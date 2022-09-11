package main

import "fmt"

func main() {
	// Long form assignment
	var mytext = "test string 1"
	// short form assigment
	mytext2 := "test string 2"

	// multiline string
	mytext3 := `long string 
spanning multiple lines`

	fmt.Println(mytext)
	fmt.Println(mytext2)
	fmt.Println(mytext3)
}
