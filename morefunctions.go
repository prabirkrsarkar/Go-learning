package main

import "fmt"

// Go-style: Check edge cases and return early if the input is invalid.
// Rercursion: A function calls itself until it reaches a base condition, then exists and returns a value back up the call stack.
func factorial(n int) int {
	if n < 0 {
		fmt.Println("Factorial is not defined for negative numbers.")
		return -1 // Return -1 to indicate an error for negative input
	}
	// Base case: factorial of 0 or 1 is 1.
	// This stops the recursion from going indefinitely which results in stack overflow.
	if n == 0 || n == 1 {
		return 1 // Base case: factorial of 0 or 1 is 1
	}
	return n * factorial(n-1) // Recursive case: n! = n * (n-1)!
}

// Closures: A function that captures and retains access to variables from its surrounding scope,
// even after that scope has finished executing.

// Here, the sequence(n int) function is the The Factory function that creates the environment
// and gives birth to a closure (the anonymous function).

// The anonymous functions is the closure because it "closes over" or captures the variable current
// from the parent scope, allowing it to maintain state across multiple calls to the returned function.

func sequence(step int) func() int {
	current := 0        // Initialize the current value of the sequence to 0
	return func() int { // Return an anonymous function that generates the next number in the sequence each time it is called
		current += step // Increment the current value by step
		return current  // Return the current value of the sequence
	}
}

func main() {
	fmt.Println(factorial(-3)) // Test the edge case where n is negative
	fmt.Println(factorial(0))
	fmt.Println(factorial(4))

	sq1 := sequence(1) // Create a new sequence generator
	sq2 := sequence(2) // Create another sequence generator with a different step

	fmt.Println(sq1()) // Output: 1
	fmt.Println(sq1()) // Output: 2
	fmt.Println(sq1()) // Output: 3

	fmt.Println(sq2()) // Output: 2
	fmt.Println(sq2()) // Output: 4
	fmt.Println(sq2()) // Output: 6
}
