package main 

import "fmt"

func pointers() {

	var Name = "Future Mulenga"

	var namePtr = &Name

	fmt.Println("Name:", Name) // Printing the value of Name
	fmt.Println("Name Pointer:", namePtr) // Printing the address of Name
	fmt.Println("Value at Name Pointer:", *namePtr) // Dereferencing the pointer to get the value of Name
	fmt.Println("Address of Name Pointer:", &namePtr) // Printing the address of the pointer itself
}
