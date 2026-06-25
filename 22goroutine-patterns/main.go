package main

import (
	// "goroutine-pattern/pipeline"
	// workerfanout "goroutine-pattern/worker-fanout"
	// donechannel "goroutine-pattern/done-channel"
	semaphore "goroutine-pattern/semaphore"
)

func main() {
	// workerfanout.RunWorkerFanout()
	// pipeline.RunPipeline()
	// donechannel.RunDoneChannel()
	semaphore.RunSemaphore()
}
