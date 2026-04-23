package main

import "fmt"

func main() {
	var a int
	fmt.Println(a) // Go assigned a default zero value of 0 to a

	var b string
	b = "Hello, Go!"
	fmt.Println(b) // Output: Hello, Go!

	var c float64 = 3.14
	fmt.Println(c) // Output: 3.14

	var isRunning bool = true
	fmt.Println(isRunning) // Output: true

	var firstName, lastName string
	firstName = "John"
	lastName = "Doe"
	fmt.Println(firstName, lastName) // Output: John Doe

	// Short variable declaration within a function
	message := "Hello, Go!" // Go compiler infers the type of the variable 'message' as string
	fmt.Println(message)    // Output: Hello, Go!

	// In Go, the short variable declaration operator (:=) cannot be used outside
	// of a function body. If you try to use it at the "package level"
	// (the global space at the top of your file), the compiler will throw an
	// error: syntax error: non-declaration statement outside function body.
	//The := operator is designed specifically for local variables to make code
	// within functions more concise and readable.

	const pi float32 = 3.14
	fmt.Println(pi) // Output: 3.14

	// pi = pi + 2
	// Uncommenting the above line will cause a compile-time error because constants
	// cannot be reassigned.
	// error: cannot assign to pi (neither addressable nor a map index expression)

}
