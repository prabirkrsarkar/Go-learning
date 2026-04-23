package main

import (
	"errors"
	"fmt"
)

// ValidationError is a custom error type with extra context
// about which field failed and why.
type ValidationError struct {
	Field   string // the name of the field that failed validation
	Message string // a human-readable explanation
}

// Error implements the error interface for ValidationError.
// Any type with an Error() string method satisfies the built-in error interface.
func (v *ValidationError) Error() string {
	return fmt.Sprintf("validation failed on field %q: %s", v.Field, v.Message)
}

// validateAge checks that age is at least 18.
// It returns a *ValidationError wrapped as an error if the check fails.
func validateAge(age int) error {
	if age < 18 {
		return &ValidationError{
			Field:   "age",
			Message: fmt.Sprintf("must be at least 18, got %d", age),
		}
	}
	return nil // no error – age is valid
}

func main() {
	// Try validating an invalid age
	err := validateAge(15)
	if err != nil {
		fmt.Println("Error:", err)

		// Use errors.As to check if the error is (or wraps) a *ValidationError.
		// This is the recommended way in Go to inspect custom error types.
		var ve *ValidationError
		if errors.As(err, &ve) {
			fmt.Println("Field  :", ve.Field)
			fmt.Println("Message:", ve.Message)
		}
	}

	// Try validating a valid age
	err = validateAge(25)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Age 25 is valid!")
	}
}
