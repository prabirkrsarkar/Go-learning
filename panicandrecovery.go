package main

import "fmt"

func mightPanic(shouldPanic bool) {
	if shouldPanic {
		panic("A panic occurred in mightPanic!")
	}

	fmt.Println("This function executed without panicking.")
}

func main() {

	// Do not use this as a normal way to handle errors in Go.
	// This is just for demonstration purposes.
	// panic("Something went wrong!")

	// Defer a function to recover from a panic.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	mightPanic(false)
	mightPanic(true) // This will cause a panic, but it will be recovered by the deferred function.
	fmt.Println("This will not be printed because the panic occurs before this line.")
}
