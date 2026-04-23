package main

import "fmt"

// Go-style: Use a variadic function to handle multiple arguments.
func sum(numbers ...int) int { // here, numbers is a slice of integers
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Optional argument via variadic parameters:
// The first argument is the name, and the second is an optional greeting message.
func greet(name string, greetings ...string) {
	msg := "Hello" // Default value for greetings set to "Hello"
	if len(greetings) > 0 {
		msg = greetings[0] // Override default if a greeting is provided
	}
	fmt.Println(msg, name)
}

func main() {
	fmt.Println("Sum :", sum(1, 3, 5, 7))
	greet("Alice")
	greet("Bob", "Hi")
}
