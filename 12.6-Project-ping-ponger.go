package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// ping runs as a separate goroutine. It keeps sending "ping" messages
// to ch until the context is cancelled. The WaitGroup allows main() to know
// when this goroutine has completely finished.
// Calling wg.Done() reduces the WaitGroup counter by one.
func ping(ctx context.Context, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for { // infinite loop
		select {
		// ctx.Done() becomes readable when cancel() is called.
		// This tells the goroutine to stop its work and return.
		case <-ctx.Done():
			fmt.Println("ping goroutine stopped")
			return
		// Try to send a ping message to the channel.
		case ch <- fmt.Sprintf("ping: %v", time.Now()):
		}
		// Wait before producing the next message.
		//
		// This sleep is outside the select, so cancellation may take
		// up to one second to be noticed.
		time.Sleep(1 * time.Second)
	}
}

// pong works in the same way as ping, but sends "pong" messages.
func pong(ctx context.Context, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for { // infinite loop
		select {
		// ctx.Done() channel becomes readable when cancel() is called.
		// This tells the goroutine to stop its work and return.
		case <-ctx.Done():
			fmt.Println("pong goroutine stopped")
			return
		// Try to send a pong message to the channel.
		case ch <- fmt.Sprintf("pong: %v", time.Now()):
		}
		// Wait before producing the next message.
		//
		// This sleep is outside the select, so cancellation may take
		// up to one second to be noticed.
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// Create a cancellable context.
	// ctx is passed to the worker goroutines.
	// Calling cancel() closes ctx.Done(), telling the workers to stop.
	ctx, cancel := context.WithCancel(context.Background())

	// Safety net:
	// if main() exits through any path, cancel the context.
	defer cancel()

	// Channel used by ping() and pong() to send messages to main().
	pingerCh := make(chan string)

	// WaitGroup tracks the two worker goroutines.
	var wg sync.WaitGroup
	wg.Add(2)

	// Start the worker goroutines.
	go ping(ctx, pingerCh, &wg)
	go pong(ctx, pingerCh, &wg)

	// After five seconds, timeout receives a value.
	timeout := time.After(5 * time.Second)

	// Keep receiving messages until the timeout occurs.
	for {
		select {
		case <-timeout:
			fmt.Println("Timeout!!")
			// Explicitly cancel now so that ping() and pong() stop.
			cancel()

			// Wait until both goroutines have returned.
			wg.Wait()

			fmt.Println("Done!")
			return

		case msg := <-pingerCh:
			// Print a message received from either ping() or pong().
			fmt.Println(msg)

		}
	}
}
