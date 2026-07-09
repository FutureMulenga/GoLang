package main 

import "fmt"


func main() {
	
	var firstName string  = "Future" // Explicit type declaration, specifying that firstName is of type string
	var lastName  = "Mulenga" // Type inference, Go will automatically determine the type based on the value assigned
	greeting := "Hello, my name is " + firstName + " " + lastName // Concatenation of strings using the + operator

	fmt.Println(greeting)
}

