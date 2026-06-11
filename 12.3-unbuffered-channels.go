package main

type user struct {
	name string
}

func main() {
	// Create channels to communicate between goroutines
	// This channel can create and receive strings
	messages := make(chan string)

	// This channel can create and receive values of type user
	users := make(chan user)

	go func() {
		// Send a message to the channel
		messages <- "Hello Prabir from a goroutine!"
		messages <- "Hello Prabir from a goroutine again!"
	}() // Immediately invoked anonymous function

	go func() {
		// Send a user struct to the users channel
		users <- user{name: "Prabir"}
	}() // Immediately invoked anonymous function

	/* The main goroutine will block until it receives a message from the channel
	   Receive a value from the messages channel.
	   If no value is available yet, the main goroutine blocks (waits).
	   Every receive must eventually have a matching send (for unbuffered channels)
	*/
	msg := <-messages
	println(msg)

	msg = <-messages
	println(msg)

	userschannel := <-users
	println("Received user:", userschannel.name)
}
