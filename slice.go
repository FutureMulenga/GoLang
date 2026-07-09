package main 

import "fmt"

func slices(){

	numbers := []int{1, 2, 3, 4, 5} // Slice of integers, which is a dynamically-sized array

	fmt.Println("Numbers:", numbers, "Length:", len(numbers)) // Printing the slice and its length using the len() function

	slice := numbers[1:4] // Creating a slice from the original array, starting from index 1 to index 4 (exclusive)
	fmt.Println("Slice:", slice, "Length:", len(slice)) // Printing the new slice and its length
}