package main

import "fmt"

func main() {
	// C-style for loop
	for i := 0; i < 5; i++ {
		fmt.Printf("C-style for loop iteration: %d\n", i)
	}

	//While-style for loop
	j := 0
	for j < 5 {
		fmt.Printf("While-style for loop iteration: %d\n", j)
		j++
	}

	// A possible, Infinite loop (use with caution)
	counter := 0
	for {
		if counter >= 5 {
			break // Exit the loop when counter reaches 5
		}
		fmt.Printf("Infinite loop iteration: %d\n", counter)
		counter++
	}

	// Using for loop to skip even numbers
	for k := 0; k < 10; k++ {
		if k%2 == 0 {
			continue // Skip even numbers
		}
		fmt.Printf("Odd number: %d\n", k)
	}

	// Iterating over an array using for loop
	numbers := []int{1, 2, 3, 4, 5}
	for index := 0; index < len(numbers); index++ {
		fmt.Printf("Array element at index %d: %d\n", index, numbers[index])
	}

	//Iterating over a string array using range
	strings := []string{"Go", "is", "awesome"}
	for index, value := range strings {
		fmt.Printf("String element at index %d: %s\n", index, value)
	}

	// Discard the index using _ when you don't need it
	for _, value := range strings {
		fmt.Printf("String element: %s\n", value)
	}
}

// In Go, the range keyword is a specialized tool used to "walk through"
// a collection of data. When you use it on a slice or an array, it returns
// two pieces of information for every iteration: the index and the value
// of the element at that index.
