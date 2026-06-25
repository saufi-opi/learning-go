package pipeline

import "fmt"

// Stage 1 — source. Takes no channel, produces one.
func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out) // close when the loop below ends
		for _, n := range nums {
			out <- n // blocks until the next stage receives
		}
	}()
	return out // returns instantly; the goroutine runs in the background
}

// Stage 2 — transform. Reads in, sends squares to out.
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out) // close out after in is drained
		for n := range in { // ends when the previous stage closes its channel
			out <- n * n
		}
	}()
	return out
}

// Stage 3 — filter. Only passes even numbers through.
func filterEven(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			if n%2 == 0 {
				out <- n
			}
		}
	}()
	return out
}

func RunPipeline() {
	// Compose the stages — reads like a sentence.
	nums := generate(1, 2, 3, 4, 5, 6, 7, 8)
	squared := square(nums)
	evens := filterEven(squared)

	// main is the final consumer. Drains the last channel.
	for n := range evens {
		fmt.Println(n)
	}
}