package main

import "fmt"

// Go-style: Check edge cases and return early if the input is invalid.
func calculateArea(length float64, width float64) float64 {
	if length <= 0 || width <= 0 {
		fmt.Println("Length and width must be positive numbers.")
		return 0 // Return 0 if either length or width is not a positive number
	}
	return length * width // Calculate and return the area of the rectangle
}

func main() {
	area := calculateArea(5, 3)
	fmt.Printf("The area of the rectangle is: %.2f\n", area) // Print the calculated area of the rectangle with 2 decimal places

	area = calculateArea(-5, 3)                              // Test the edge case where length is negative
	fmt.Printf("The area of the rectangle is: %.2f\n", area) // Print the result, which should be 0 due to invalid input

	// Go-style: Assign an anonymous function to a variable.
	logger := func(message string) {
		fmt.Println("Log:", message) // Log the provided message with a "Log:" prefix
	}

	logger("This is a log message.") // Call the logger function with a sample message

}
