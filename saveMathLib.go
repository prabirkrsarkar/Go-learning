package main

import (
	"fmt"
	"strings"
)

type MathError struct {
	Operation string
	InputA    int
	InputB    int
	Message   string
}

// Error implements the error interface for MathError.
func (e MathError) Error() string {
	var inputs []string
	if e.Operation == "Division" {
		inputs = []string{
			fmt.Sprintf("InputA: %d", e.InputA),
			fmt.Sprintf("InputB: %d", e.InputB),
		}
	}
	// Using append is Most common
	// Works with unknown size
	// Handles capacity automatically

	// if e.Operation == "Division" {
	//	inputs = append(inputs, fmt.Sprintf("InputA: %d", e.InputA))
	//	inputs = append(inputs, fmt.Sprintf("InputB: %d", e.InputB))

	return fmt.Sprintf("Operation: %s, %s, Message: %s", e.Operation, strings.Join(inputs, ", "), e.Message)
}

func safeDivison(a, b int) (int, error) {
	if b == 0 {
		return 0, MathError{
			Operation: "Division",
			InputA:    a,
			InputB:    b,
			Message:   "division by zero",
		}
	}
	return a / b, nil
}

// safeDivison() returns a MathError
// Go sees MathError implements error → wraps it as error
// err now holds a value of type error (backed by MathError)
// fmt.Println sees an error
// It calls Error() automatically
// Your formatted string is printed
func main() {
	result, err := safeDivison(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
