package main

import (
	"fmt"
	"slices"
)

// While slicing include low index, exclude high index
func main() {
	fmt.Println("---Advanced Slicing Operations---")
	original := []int{1, 2, 3, 4, 5}
	fmt.Println("Original slice, len, cap):", original, len(original), cap(original))

	// Slicing with a different range. While slicing include low index, exclude high index
	slice1 := original[1:4] // This will include elements at index 1, 2, and 3
	fmt.Println("Slice1 (original[1:4], len, cap):", slice1, len(slice1), cap(slice1))

	// Slicing with a different range and capacity. While slicing include low index, exclude high index
	slice2 := original[2:5] // This will include elements at index 2, 3, and 4
	fmt.Println("Slice2 (original[2:5], len, cap):", slice2, len(slice2), cap(slice2))

	slice3 := original[:4] // This will include elements at index 0, 1, 2, and 3
	fmt.Println("Slice3 (original[:4], len, cap):", slice3, len(slice3), cap(slice3))

	slice4 := original[2:] // This will include elements at index 2, 3, and 4. So, here, includes low index and goes till the end
	fmt.Println("Slice4 (original[2:], len, cap):", slice4, len(slice4), cap(slice4))

	// slice5 is just a clone of original slice
	slice5 := original[:] // This will include all elements from the original slice
	fmt.Println("Slice5 (original[:], len, cap):", slice5, len(slice5), cap(slice5))

	// Using the slices package and it's functions.
	fmt.Println("Does slice4 contain the element 3?", slices.Contains(slice4, 3))

	slice6 := append(original, 6, 7, 8) // Appending elements to the original slice
	fmt.Println("Slice6 (after appending 6, 7, 8):", slice6, len(slice6), cap(slice6))
	fmt.Println("Original slice after appending:", original, len(original), cap(original))

}
