package main

import (
	"fmt"
	"time"
)

// CHANNEL - Komunikasi antar goroutine
func main() {
	fmt.Println("=== CHANNEL EXAMPLES ===")

	// Channel dasar
	ch := make(chan string)

	go func() {
		ch <- "Hello from goroutine"
	}()

	msg := <-ch
	fmt.Println("Received:", msg)

	// Channel dengan loop
	dataCh := make(chan int, 5)

	go func() {
		for i := 1; i <= 5; i++ {
			dataCh <- i
			time.Sleep(100 * time.Millisecond)
		}
		close(dataCh) // Tutup channel setelah selesai
	}()

	// Baca dari channel yang sudah di-close
	for val := range dataCh {
		fmt.Println("Received:", val)
	}

	// Bidirectional channel
	responseCh := make(chan string)

	go processRequest("GET /api/users", responseCh)

	response := <-responseCh
	fmt.Println("Response:", response)
}

func processRequest(req string, response chan<- string) {
	time.Sleep(200 * time.Millisecond)
	response <- fmt.Sprintf("Processed: %s", req)
}
