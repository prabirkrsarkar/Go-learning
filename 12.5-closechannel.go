package main

import (
	"fmt"
	"sync"
)

func main() {

	jobs := make(chan int, 5)
	//done := make(chan bool)
	//We can use a Waitgroup instead of this second 'done' channel

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		// Reading messages from jobs from this goroutine
		for {
			// Comma, OK patterm; as we seen in map, we get the message in the channel
			// ok returns false if the channel is closed.
			r, ok := <-jobs

			if ok {
				fmt.Println("Got this messsage", r)
			} else {
				fmt.Println("Channel is closed")
				return
			}
		}
	}()

	// Sending messages into jobs channel from Main goroutine

	for i := 1; i <= 4; i++ {
		jobs <- i
		fmt.Println("Sending message", i)
	}

	close(jobs)

	// This receive will make the main Go routine wait for sender.
	// If there are no senders then program will fail
	// Alternative to Waitgroup
	wg.Wait()
}
