package main

import "fmt"

func main() {
	// Creating and initializing a
	// variable with string literal
	// using double quotes
	string_value_1 := "Welcome to test golang"

	//adding escape character
	string_value_2 := "Weclome!\n golang"
	//Using backtickst
	string_value_3 := `Hello!allforchecking on the golang`

	string_value_4 := `Hello!\n testfortest`

	//Display strings
	fmt.Println("string 1: ", string_value_1)
	fmt.Println("String 2: ", string_value_2)
	fmt.Println("String 3: ", string_value_3)
	fmt.Println("String 4: ", string_value_4)
}
