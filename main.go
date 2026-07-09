package main 

import "fmt"


func dataTypes() {

	var age int8 = 25 // Explicit type declaration, specifying that age is of type int
	var uintAge uint8 = 30 // Explicit type declaration, specifying that uintAge is of type uint
    var height float32 = 5.9 // Explicit type declaration, specifying that height is of type float32


	fmt.Println("Age:", age)
	fmt.Println("Unsigned Age:", uintAge)
	fmt.Println("Height:", height)
}


func main() {
	
	dataTypes()
	var firstName string  = "Future" // Explicit type declaration, specifying that firstName is of type string
	var lastName  = "Mulenga" // Type inference, Go will automatically determine the type based on the value assigned
	greeting := "Hello, my name is " + firstName + " " + lastName // Concatenation of strings using the + operator

	fmt.Println(greeting)
}

