package main

import (
	"fmt"
	"time"
)

// GOROUTINE - Contoh dasar
func main() {
	fmt.Println("=== GOROUTINE EXAMPLES ===")

	// Goroutine sederhana
	go func() {
		fmt.Println("Goroutine 1: Running in background")
	}()

	// Goroutine dengan parameter
	go printMessage("Goroutine 2", 3)

	// Goroutine dengan multiple workers
	for i := 1; i <= 5; i++ {
		go func(id int) {
			fmt.Printf("Worker %d is working\n", id)
		}(i)
	}

	time.Sleep(1 * time.Second) // Tunggu goroutine selesai
}

func printMessage(msg string, times int) {
	for i := 0; i < times; i++ {
		fmt.Printf("%s: %d\n", msg, i+1)
		time.Sleep(100 * time.Millisecond)
	}
}
