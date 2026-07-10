package main 

import "fmt"

func main() {

	maps()

	// initial1, initial2 := listNames("Future Mulenga")
	// fmt.Println("First Name Initial:", initial1)
	// fmt.Println("Last Name Initial:", initial2)
	// loops()
	// slices()
	// arrays()

	// formatStrings()
	// dataTypes()

	var firstName string  = "Future" // Explicit type declaration, specifying that firstName is of type string
	var lastName  = "Mulenga" // Type inference, Go will automatically determine the type based on the value assigned
	greeting := "Hello, my name is " + firstName + " " + lastName // Concatenation of strings using the + operator

	fmt.Println(greeting)
}

