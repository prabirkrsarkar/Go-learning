package main

import (
	"errors"
	"fmt"
	"strings"
)

var ErrDivisionByZero = errors.New("Division by zero is not allowed")
var ErrNumberTooLarge = errors.New("Number is too large")
var ErrInvalidFullName = errors.New("Invalid full name, expected at least first and last name")

// Go-Style: Check for errors and return early if error occurs.
// This is a common pattern in Go, where functions return multiple values, including an error.
// By default error is the last return value, and it is common to check for errors immediately after calling a function.
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	if a > 10000 {
		return 0, ErrNumberTooLarge
	}
	return a / b, nil
}

// Named return values: Go allows you to name the return values in a function signature.
// This can make the code more readable and allows you to return values
// without explicitly specifying them in the return statement.
func splitName(fullName string) (firstName, lastName string, err error) {
	parts := strings.Split(fullName, " ")
	if len(parts) < 2 {
		return "", "", ErrInvalidFullName
	}
	firstName = parts[0]
	lastName = parts[1]
	return // Note the empty return statement, which returns the named return values
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		if errors.Is(err, ErrDivisionByZero) {
			fmt.Println("DividebyZero!:", err)
		} else if errors.Is(err, ErrNumberTooLarge) {
			fmt.Println("NumberTooLarge!:", err)
		} else {
			fmt.Println("Unexpected error:", err)
		}

		fmt.Println("Result:", result)
	}

	firstName, lastName, err := splitName("Prabir")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("First Name: %s, Last Name: %s\n", firstName, lastName)
	}
}
