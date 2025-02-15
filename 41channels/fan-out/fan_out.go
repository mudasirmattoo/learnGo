/*
How Fan-Out Works
Single Input Channel:
A single channel provides tasks or data to be processed.

Multiple Worker Goroutines:
Multiple goroutines (workers) read from this channel and process the data in parallel.

Parallel Processing:
Each worker picks up data independently, allowing multiple computations to happen simultaneously.*/

/*
In Go, fan-out is a concurrency pattern where you distribute work from a single channel to multiple worker goroutines. This pattern is useful when you have a single source of data (e.g., a channel) and want to process it in parallel using multiple workers
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker function that processes data from a channel
func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Second)
	}
}
func main() {
	const numOfWorkers = 3 //NUMBER OF GOROUTINES YOU WANT
	const numOfJobs = 10

	// Create a channel for jobs
	jobs := make(chan int, numOfJobs)

	// Create a WaitGroup to wait for all workers to finish
	var wg sync.WaitGroup

	//start worker go-routines
	for i := 1; i <= numOfWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	//send jobs to the channel

	for k := 1; k <= numOfJobs; k++ {
		jobs <- k
	}

	close(jobs) // Close the channel to signal no more jobs

	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("about to quit")

	/*
				OUTPUT
		-------------------------
		Worker 2 processing job 2
		Worker 1 processing job 1
		Worker 3 processing job 3
		Worker 3 processing job 4
		Worker 2 processing job 6
		Worker 1 processing job 5
		Worker 1 processing job 7
		Worker 3 processing job 8
		Worker 2 processing job 9
		Worker 1 processing job 10
		about to quit
	*/
}
