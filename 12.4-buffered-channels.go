package main

import "fmt"

func main() {

	messages := make(chan string, 3)

	fmt.Println("Sending message to a buffered channel")
	messages <- "Message 1 into channel"
	messages <- "Message 2 into channel"
	messages <- "Messages 3 into channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	// At this point, the channel is empty.

	// The fourth receive waits for another value:
	// But there is no other goroutine that can send a value, so Go detects:
	// fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<-messages)

}
