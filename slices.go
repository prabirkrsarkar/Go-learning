package main

import "fmt"

func main() {
	names := []string{"Alice", "Bob", "Charlie"} // A Slice, note size is not mentioned, Memory is allocated.
	fmt.Println(names)

	// names_arr := [3]string{"Alice", "Bob", "Charlie"} // An Array since size is mentioned.

	strings := []string{} // An Empty slice but NOT a Nil slice
	fmt.Println("Empty but NOT Nil Slice:", strings)

	// var morenames []string // A Nil slice, since declaring a slice doesn't allocate memory for its elements.
	// fmt.Println(morenames[0]) // panic: runtime error: index out of range [0] with length 0

	items := make([]int, 3, 5) // Allocates memory for an underlying array of capacity 5
	fmt.Println(items)         // prints [0 0 0] since the slice is initialized with zero values
	// fmt.Println(items[3])  // panic: runtime error: index out of range [3] with length 3 since

	items = append(items, 1)
	items = append(items, 2)
	items = append(items, 3) // This will cause the slice to grow beyond its initial capacity, creating a new underlying array.

	fmt.Println(len(items)) // returns the current length of the slice, which is 6 after appending three elements
	fmt.Println(cap(items)) // returns the current capacity of the slice, which is doubled of initial capacity (10) after appending elements beyond the initial capacity

}
