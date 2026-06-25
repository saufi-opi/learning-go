package donechannel

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func RunDoneChannel() {
	// can use a context to signal done to multiple goroutines at once or just a channel. Both are valid approaches.
	ctx, cancel := context.WithCancel(context.Background())
	// done := make(chan struct{})

	var wg sync.WaitGroup
	for i := range 3 {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("worker %d: stopping\n", id)
					return
				default:
					fmt.Printf("worker %d: working\n", id)
					time.Sleep(300 * time.Millisecond)
				}
			}
		}(i)
	}

	time.Sleep(1 * time.Second)
	cancel() // broadcast stop to ALL goroutines at once
	// close(done) // broadcast stop to ALL goroutines at once
	
	wg.Wait()
	fmt.Println("all done")
}
