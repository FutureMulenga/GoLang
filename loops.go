package main 

import "fmt"


func loops() {

	x := 0

	for x < 5 { // Using a for loop to iterate while x is less than 5
		fmt.Println("Value of x:", x) // Printing the current value of x
		x++ // Incrementing x by 1
	}


	names := []string{"Alice", "Bob", "Charlie", "David", "Eve"} // Slice of strings containing names

	for i := 0; i < len(names); i++ { // Using a for loop to iterate over the slice of names
		fmt.Println("Name:", names[i]) // Printing each name in the slice
	}

}