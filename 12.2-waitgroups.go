package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello(message string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(delay)
	fmt.Println("From sayHello:", message)
}

func main() {

	/* Rules for using wait groups
	1. Define the wait group outside of the goroutines.
	2. Add to the wait group counter before starting each goroutine with wg.Add().
	3. Decrement the wait group counter when a goroutine completes with wg.Done().
	4. Call `Wait` on the wait group to block until all goroutines have finished with wg.Wait().
	5. Always pass a pointer to the wait group, not a copy,to the goroutines so they can modify the same wait group instance.
	*/

	// Create a wait group to synchronize the completion of goroutines
	var wg sync.WaitGroup

	totalJobs := 3

	// The goruotines can complete in any order, but the main goroutine will wait for all of them
	// to finish before exiting.
	for i := 0; i < totalJobs; i++ {
		wg.Add(1) // Increment the wait group counter for each job
		go sayHello(fmt.Sprintf("Hello from a Job %d, goroutine!", i), time.Second, &wg)
	}

	fmt.Println("Hello from Main() go routines!")

	wg.Wait() // Wait for all goroutines to finish

}
