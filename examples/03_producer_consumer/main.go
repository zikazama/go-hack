package main

import (
	"fmt"
	"time"
)

// PRODUCER-CONSUMER PATTERN
func main() {
	fmt.Println("=== PRODUCER-CONSUMER PATTERN ===")

	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// Producer - menghasilkan pekerjaan
	go producer(jobs)

	// Consumer - mengolah pekerjaan (3 workers)
	for w := 1; w <= 3; w++ {
		go consumer(w, jobs, results)
	}

	// Kumpulkan hasil
	for i := 1; i <= 10; i++ {
		result := <-results
		fmt.Printf("Result: %d\n", result)
	}
}

func producer(jobs chan<- int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("Producer: Sending job %d\n", i)
		jobs <- i
		time.Sleep(200 * time.Millisecond)
	}
	close(jobs)
}

func consumer(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Consumer (Worker %d): Processing job %d\n", id, job)
		time.Sleep(500 * time.Millisecond)
		results <- job * 2
	}
}
