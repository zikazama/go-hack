package main

import (
	"fmt"
	"time"
)

// WORKER POOL PATTERN
func main() {
	fmt.Println("=== WORKER POOL PATTERN ===")

	const numWorkers = 3
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs
	numJobs := 9
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	for r := 1; r <= numJobs; r++ {
		result := <-results
		fmt.Printf("Collected result: %d\n", result)
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d: Started job %d\n", id, job)
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("Worker %d: Finished job %d\n", id, job)
		results <- job * 2
	}
}
