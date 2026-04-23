package main

import "fmt"

// In Go, arguments are passed by value. This means that when you pass a
// variable to a function, the function receives a copy of the variable,
// not the original variable itself.

// But when you pass a pointer to a function, the function can modify
// the original variable.

func modifyInt(n int) {
	fmt.Println("Original value of n passed to this function is: ", n)
	n = n * 100
	fmt.Println("Modified value of n:", n)
}

// Note we are using the Go's idiomatic Guard clauses, aka Early return
// i.e. getting rid of the corner cases (null pointer) first
func modifyPointer(p *int) { // Argument to this function is a pointer to an int
	// Guard clauses handle errors/nil first and return early
	if p == nil {
		fmt.Println("A Null Pointer!")
		return
	}

	// After the guard clause, we are 100% sure p is safe to use
	fmt.Println("Original value at p (*p) passed to this function is: ", *p)
	fmt.Println("Modifying value at pointer:", p)
	*p = *p * 100 // Dereferencing the pointer to modify the value it points to
	fmt.Println("Modified value at the pointer p:", *p)
}

func main() {

	age := 45
	agePointer := &age // The & operator is used to get the memory address of the variable age
	fmt.Println("Age:", age)
	fmt.Println("Age Pointer:", agePointer)
	fmt.Println("Age Pointer:", &age) // This will give the same memory address as agePointer

	fmt.Println("Value at Age Pointer:", *agePointer)
	fmt.Println("Value at Age Pointer written as *(&age) is: ", *(&age))

	modifyInt(age)
	fmt.Println("Age after modifyInt:", age) // Original value remains unchanged

	modifyPointer(&age)
	fmt.Println("Age after modifyPointer:", age) // Value has been modified through the pointer

}
