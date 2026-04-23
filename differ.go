package main

import (
	"fmt"
	"os"
)

// Defer is a powerful feature in Go that allows you to
// schedule a function call to be executed after the surrounding function returns.
// The deferred function will run even if the surrounding function panics, making it useful for cleanup tasks like closing files or releasing resources.
func simpleDefer() {
	fmt.Println("Function simpleDefer: started")
	defer fmt.Println("Function simpleDefer: deferred statement executed")
	fmt.Println("Function simpleDefer in progress")
	fmt.Println("Function simpleDefer: about to return")
	fmt.Println("Function simpleDefer: returning")
}

// Deferred functions are executed in LIFO (Last In, First Out) order,
// meaning that if you defer multiple functions, they will be executed in reverse order of their defer statements.
func lifomultipleDefers() {
	fmt.Println("Function lifomultipleDefers: started")
	defer fmt.Println("First deferred statement executed")
	defer fmt.Println("Second deferred statement executed")
	defer fmt.Println("Third deferred statement executed")
	fmt.Println("Function lifomultipleDefers: about to return")
	fmt.Println("Function lifomultipleDefers: returning")
}

func main() {
	simpleDefer()
	lifomultipleDefers()

	file, err := os.Create("./defer_example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close() // Close the file when main() returns, ensuring it happens even if there's an error later on.
}
