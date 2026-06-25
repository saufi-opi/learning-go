package workerfanout

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type result struct {
	url    string
	status int
	err    error
}

// checkURL does the actual work. This part is done for you —
// the concurrency is the point, not the HTTP call.
func checkURL(url string) result {
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return result{url: url, err: err}
	}
	defer resp.Body.Close()
	return result{url: url, status: resp.StatusCode}
}

// worker reads URLs from `jobs`, checks each, sends to `results`.
func worker(jobs <-chan string, results chan<- result, wg *sync.WaitGroup) {
	// TODO 1: signal done when this worker finishes (defer ...)
	defer wg.Done()
	// TODO 2: range over jobs; for each url, send checkURL(url) into results
	for url := range jobs {
		results <- checkURL(url)
	}
}

func RunWorkerFanout() {
	urls := []string{
		"https://golang.org",
		"https://github.com",
		"https://httpstat.us/404",
		"https://httpstat.us/500",
		"https://this-domain-does-not-exist-xyz.com",
		"https://pkg.go.dev",
	}

	const numWorkers = 3

	jobs := make(chan string)
	results := make(chan result)
	var wg sync.WaitGroup

	// TODO 3: start numWorkers workers (remember wg.Add)
	wg.Add(numWorkers)
	for range numWorkers {
		go worker(jobs, results, &wg)
	}

	// TODO 4: feed all urls into `jobs` in a goroutine, then close(jobs).
	//         Why a goroutine? Think about what blocks if you don't.
	go func() {
		for _, url := range urls {
			jobs <- url // blocks until a worker receives this one
		}
		close(jobs) // reached only after all sends complete
	}()

	// TODO 5: close(results) once all workers are done.
	//         Which of the four blocking rules tells you where this goes?
	go func() {
		wg.Wait()
		close(results)
	}()

	// TODO 6: range over results and print each one.
	for res := range results {
		if res.err != nil {
			fmt.Printf("%s: error: %v\n", res.url, res.err)
		} else {
			fmt.Printf("%s: status code %d\n", res.url, res.status)
		}
	}
}
