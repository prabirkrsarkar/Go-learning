package main

import "fmt"

// This code demonstrates the use of variadic functions in Go, which allow you to pass a variable number of arguments to a function. The Sum function takes an integer 'a' and a variadic parameter '...int', which can accept any number of integers. The function calculates the total sum of all the integers
// passed as arguments and returns the result. But this would not work with floating point numbers,
// so we can use generics to make it work with both integers and floating point numbers.

func sum(a ...int) int {
	total := 0
	for _, num := range a {
		total += num
	}
	return total
}

// T is the type parameter, and it can be any type that
// satisfies the constraint specified in the square brackets.
// In this case, we are using a type constraint that allows T to be either int or float64.
// The function takes a variable number of arguments of type T and returns a value of type T,
// which is the sum of all the arguments passed to the function.

func addGenerics[T int | float64 | string](a ...T) T {
	total := T(0) // We need to initialize total to the zero value of type T, which is done by converting 0 to type T.
	for _, num := range a {
		total += num
	}
	return total
}

// Defining a type constraint interface named addables that allows T
// to be either int, float64, or string.
type addables interface {
	int | float64 | string
}

func addGenericsWithConstraint[T addables](a ...T) T {
	total := T(0) // We need to initialize total to the zero value of type T, which is done by converting 0 to type T.
	for _, num := range a {
		total += num
	}
	return total
}

func main() {
	// Using the sum function with integers
	fmt.Println("Sum of integers:", sum(1, 2, 3, 4, 5))

	// Using the sumGenerics function with integers
	fmt.Println("Sum of integers using generics:", addGenerics(1, 2, 3, 4, 5))

	// Using the sumGenerics function with floating point numbers
	fmt.Println("Sum of floats using generics:", addGenerics(1.5, 2.5, 3.6))

	// Using the sumGenerics function with strings
	fmt.Println("Concatenation of strings using generics:", addGenerics("Hello, ", "world!"))

	// Using the sumGenericsWithConstraint function with integers
	fmt.Println("Sum of integers using generics with constraint:", addGenericsWithConstraint(1, 2, 3, 4, 5))

	// Using the sumGenericsWithConstraint function with floating point numbers
	fmt.Println("Sum of floats using generics with constraint:", addGenericsWithConstraint(1.5, 2.5, 3.6))

	// Using the sumGenericsWithConstraint function with strings
	fmt.Println("Concatenation of strings using generics with constraint:", addGenericsWithConstraint("Hello, ", "world!"))

}
