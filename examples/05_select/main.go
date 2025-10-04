package main

import (
	"fmt"
	"time"
)

// SELECT - Memilih dari beberapa channel
func main() {
	fmt.Println("=== SELECT EXAMPLES ===")

	basicSelect()
	selectWithTimeout()
	selectWithDefault()
}

func basicSelect() {
	fmt.Println("\n--- Basic Select ---")
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutine 1
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Data from channel 1"
	}()

	// Goroutine 2
	go func() {
		time.Sleep(500 * time.Millisecond)
		ch2 <- "Data from channel 2"
	}()

	// Select - terima dari channel yang siap duluan
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Select received:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Select received:", msg2)
		}
	}
}

func selectWithTimeout() {
	fmt.Println("\n--- Select with Timeout ---")
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "This will timeout"
	}()

	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout! No message received within 1 second")
	}
}

func selectWithDefault() {
	fmt.Println("\n--- Select with Default (Non-blocking) ---")
	ch := make(chan string)

	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	default:
		fmt.Println("No message available, continuing immediately...")
	}

	// Contoh dengan send
	select {
	case ch <- "try to send":
		fmt.Println("Message sent")
	default:
		fmt.Println("Channel not ready, skipping send")
	}
}
