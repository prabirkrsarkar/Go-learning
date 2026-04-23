package main

import "fmt"

func mightPanic(shouldPanic bool) {
	if shouldPanic {
		panic("A panic occurred in mightPanic!")
	}
	fmt.Println("This function executed without panicking.")
}

// If you want execution to continue after a panic,
// recover must happen in a different function than the caller
// you want to continue
// Recover in main → main stops
// Recover in wrapper → caller (main) continues

// This functions wraps a hardcoded call to mightPanic(true)
// and recovers from the panic that occurs within it.
func recoverablePanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic in recoverablePanic:", r)
		}
	}()
	mightPanic(true)
}

// Wrapper that safely executes any function that is passed to it,
// recovering from any panics that occur within that function.
func safeExecute(fn func()) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	fn()
}

func main() {

	recoverablePanic() // This will cause a panic, but it will be recovered by the deferred function inside recoverablePanic.
	fmt.Println("This WILL be printed because panic was handled inside recoverable() NOT in main.")

	//safeExecute(func() { mightPanic(false) })
	//safeExecute(func() { mightPanic(true) })
	//fmt.Println("This WILL be printed because panic was handled inside safeExecute() NOT in main."))

}
