package main

import (
	"fmt"
	"time"
)

// BUFFERED vs UNBUFFERED CHANNEL
func main() {
	fmt.Println("=== BUFFERED vs UNBUFFERED ===")

	unbufferedExample()
	bufferedExample()
}

func unbufferedExample() {
	fmt.Println("\n--- Unbuffered Channel ---")
	// Unbuffered - blocking sampai ada receiver
	unbuffered := make(chan string)

	go func() {
		fmt.Println("Unbuffered: Sending...")
		unbuffered <- "message" // Block sampai ada yang terima
		fmt.Println("Unbuffered: Sent!")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Unbuffered: Receiving...")
	msg := <-unbuffered
	fmt.Printf("Unbuffered: Received '%s'\n", msg)
}

func bufferedExample() {
	fmt.Println("\n--- Buffered Channel ---")
	// Buffered - tidak blocking selama buffer belum penuh
	buffered := make(chan string, 2)

	go func() {
		fmt.Println("Buffered: Sending 1...")
		buffered <- "message 1" // Tidak block
		fmt.Println("Buffered: Sent 1!")

		fmt.Println("Buffered: Sending 2...")
		buffered <- "message 2" // Tidak block
		fmt.Println("Buffered: Sent 2!")

		fmt.Println("Buffered: All messages sent without blocking!")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Buffered: Now receiving...")
	fmt.Printf("Buffered: Received '%s'\n", <-buffered)
	fmt.Printf("Buffered: Received '%s'\n", <-buffered)

	// Demonstrasi buffer penuh
	fmt.Println("\n--- Buffer Full Demo ---")
	bufCh := make(chan int, 3)

	bufCh <- 1
	bufCh <- 2
	bufCh <- 3
	fmt.Println("Buffer filled with 3 items (buffer size: 3)")

	// bufCh <- 4 // Ini akan block karena buffer penuh

	fmt.Println("Reading from buffer:", <-bufCh, <-bufCh, <-bufCh)
}
