package main

import (
	"fmt"
	"strings"
)

// Go-Style: Check for errors and return early if error occurs.
// This is a common pattern in Go, where functions return multiple values, including an error.
// By default error is the last return value, and it is common to check for errors immediately after calling a function.
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}
	return a / b, nil
}

// Named return values: Go allows you to name the return values in a function signature.
// This can make the code more readable and allows you to return values
// without explicitly specifying them in the return statement.
func splitName(fullName string) (firstName, lastName string) {
	parts := strings.Split(fullName, " ")
	if len(parts) < 2 {
		fmt.Println("Invalid full name")
		return "", ""
	}
	firstName = parts[0]
	lastName = parts[1]
	return // Note the empty return statement, which returns the named return values
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	firstName, lastName := splitName("Prabir Kumar Sarkar")
	fmt.Printf("First Name: %s, Last Name: %s\n", firstName, lastName)

	firstName, lastName = splitName("Pele")
	fmt.Printf("First Name: %s, Last Name: %s\n", firstName, lastName)
}
