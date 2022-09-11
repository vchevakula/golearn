package main

import (
	"fmt"
)

func main() {
	for index, s := range `test on go lang` {
		fmt.Printf("The index number of %c is %d\n", s, index)
	}
	teststr := "check ob bytes of a string"
	// accessing the bytes of the given string
	for i := 0; i < len(teststr); i++ {
		fmt.Printf("Chracter at %d index position = %c bytevalue = %d\n", i, teststr[i], byte(teststr[i]))
	}
}
