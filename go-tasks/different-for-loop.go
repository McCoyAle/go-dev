package main

import "fmt"

func main() {
	// Basic for loop delaring, initializing and then incrementing until max num.
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Range-based for loop, setting attributes and using a range to loop through and print the index with its associated color
	colors := []string{"red", "green", "blue"}
	for index, color := range colors {
		fmt.Printf("Color at index %d is %s\n", index, color)
	}

	// Infinite loop with break statement
	count := 0
	for {
		fmt.Println("Inside infinite loop")
		count++
		if count == 5 {
			break
		}
	}
}
