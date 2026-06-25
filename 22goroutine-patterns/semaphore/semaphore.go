package semaphore

import "sync"

func RunSemaphore() {
	// Create a semaphore channel with a capacity of 3
	semaphore := make(chan struct{}, 3)

	// Function to simulate work
	work := func(id int) {
		println("Worker", id, "is working")
	}

	wg := sync.WaitGroup{}
	// Start workers
	for i := range 1000000 {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			semaphore <- struct{}{}        // Acquire the semaphore
			defer func() { <-semaphore }() // Release the semaphore

			work(id)
		}(i)
	}
	wg.Wait()
}
