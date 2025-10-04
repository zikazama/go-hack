package main

import (
	"fmt"
	"sync"
)

// FAN-OUT, FAN-IN PATTERN
func main() {
	fmt.Println("=== FAN-OUT, FAN-IN PATTERN ===")

	// Input channel
	input := make(chan int)

	// Fan-out: Buat multiple workers
	const numWorkers = 3
	channels := make([]<-chan int, numWorkers)

	for i := 0; i < numWorkers; i++ {
		channels[i] = process(i+1, input)
	}

	// Fan-in: Gabungkan hasil dari semua workers
	output := merge(channels...)

	// Send data
	go func() {
		for i := 1; i <= 10; i++ {
			input <- i
		}
		close(input)
	}()

	// Receive results
	for result := range output {
		fmt.Printf("Final result: %d\n", result)
	}
}

func process(workerID int, input <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for num := range input {
			fmt.Printf("Worker %d: Processing %d\n", workerID, num)
			output <- num * num
		}
	}()
	return output
}

func merge(channels ...<-chan int) <-chan int {
	output := make(chan int)
	var wg sync.WaitGroup

	// Start goroutine untuk setiap channel
	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for val := range c {
				output <- val
			}
		}(ch)
	}

	// Close output channel setelah semua input selesai
	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}
