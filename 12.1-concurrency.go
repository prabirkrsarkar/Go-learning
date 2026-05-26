package main

import (
	"fmt"
	"time"
)

func sayHello(message string, delay time.Duration) {
	time.Sleep(delay)
	fmt.Println("From sayHello:", message)
}

func main() {
	fmt.Println("Hello from Main() go routines!")

	go sayHello("Hello from a sayHello goroutine!", time.Second)
	go sayHello("Hello from a sayHello goroutine from 2 seconds!", 2*time.Second)
	go sayHello("Hello from a sayHello goroutine from 3 seconds!", 3*time.Second)

	fmt.Println("Last message from Main) go routines!")

	// Sleep for a while to allow the sayHello() goroutine to finish before the main goroutine exits
	// Sometime even after sleep, the goroutines may not all have completed, and the main goroutine may exit
	// before the message from sayHello goroutine is printed.

	// In a real application, you would typically use synchronization mechanisms like WaitGroups
	// to ensure that all goroutines have completed before exiting the main function.
	time.Sleep(2 * time.Second)

}
