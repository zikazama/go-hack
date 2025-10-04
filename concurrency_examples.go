package main

import (
	"fmt"
	"time"
)

// 1. GOROUTINE - Contoh dasar
func basicGoroutine() {
	fmt.Println("\n=== GOROUTINE EXAMPLES ===")

	// Goroutine sederhana
	go func() {
		fmt.Println("Goroutine 1: Running in background")
	}()

	// Goroutine dengan parameter
	go printMessage("Goroutine 2", 3)

	time.Sleep(1 * time.Second) // Tunggu goroutine selesai
}

func printMessage(msg string, times int) {
	for i := 0; i < times; i++ {
		fmt.Printf("%s: %d\n", msg, i+1)
		time.Sleep(100 * time.Millisecond)
	}
}

// 2. CHANNEL - Unbuffered & Buffered
func channelExamples() {
	fmt.Println("\n=== CHANNEL EXAMPLES ===")

	// Unbuffered channel
	unbufferedCh := make(chan string)
	go func() {
		unbufferedCh <- "Data from unbuffered channel"
	}()
	msg := <-unbufferedCh
	fmt.Println("Unbuffered:", msg)

	// Buffered channel
	bufferedCh := make(chan int, 3) // Buffer size 3
	bufferedCh <- 1
	bufferedCh <- 2
	bufferedCh <- 3
	// bufferedCh <- 4 // Akan block karena buffer penuh

	fmt.Println("Buffered:", <-bufferedCh, <-bufferedCh, <-bufferedCh)

	// Channel dengan close
	dataCh := make(chan int, 5)
	go func() {
		for i := 1; i <= 5; i++ {
			dataCh <- i
		}
		close(dataCh) // Tutup channel setelah selesai
	}()

	// Baca dari channel yang sudah di-close
	for val := range dataCh {
		fmt.Println("Received:", val)
	}
}

// 3. PRODUCER-CONSUMER PATTERN
func producerConsumer() {
	fmt.Println("\n=== PRODUCER-CONSUMER PATTERN ===")

	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// Producer - menghasilkan pekerjaan
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("Producer: Sending job %d\n", i)
			jobs <- i
			time.Sleep(200 * time.Millisecond)
		}
		close(jobs)
	}()

	// Consumer - mengolah pekerjaan (3 workers)
	for w := 1; w <= 3; w++ {
		go func(workerID int) {
			for job := range jobs {
				fmt.Printf("Consumer (Worker %d): Processing job %d\n", workerID, job)
				time.Sleep(500 * time.Millisecond)
				results <- job * 2
			}
		}(w)
	}

	// Kumpulkan hasil
	go func() {
		time.Sleep(3 * time.Second)
		close(results)
	}()

	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}
}

// 4. BUFFERED vs UNBUFFERED
func bufferedVsUnbuffered() {
	fmt.Println("\n=== BUFFERED vs UNBUFFERED ===")

	// Unbuffered - blocking sampai ada receiver
	fmt.Println("\nUnbuffered Channel:")
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

	// Buffered - tidak blocking selama buffer belum penuh
	fmt.Println("\nBuffered Channel:")
	buffered := make(chan string, 2)

	go func() {
		fmt.Println("Buffered: Sending 1...")
		buffered <- "message 1" // Tidak block
		fmt.Println("Buffered: Sent 1!")

		fmt.Println("Buffered: Sending 2...")
		buffered <- "message 2" // Tidak block
		fmt.Println("Buffered: Sent 2!")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Buffered: Receiving...")
	fmt.Printf("Buffered: Received '%s'\n", <-buffered)
	fmt.Printf("Buffered: Received '%s'\n", <-buffered)
}

// 5. SELECT - Memilih dari beberapa channel
func selectExample() {
	fmt.Println("\n=== SELECT EXAMPLES ===")

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

	// Select dengan timeout
	fmt.Println("\nSelect with timeout:")
	ch3 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch3 <- "This will timeout"
	}()

	select {
	case msg := <-ch3:
		fmt.Println("Received:", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout! No message received")
	}

	// Select dengan default (non-blocking)
	fmt.Println("\nSelect with default:")
	ch4 := make(chan string)

	select {
	case msg := <-ch4:
		fmt.Println("Received:", msg)
	default:
		fmt.Println("No message available, continuing...")
	}
}

// 6. ADVANCED: Worker Pool Pattern
func workerPool() {
	fmt.Println("\n=== WORKER POOL PATTERN ===")

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
		fmt.Printf("Worker %d: Processing job %d\n", id, job)
		time.Sleep(500 * time.Millisecond)
		results <- job * 2
	}
}

// 7. ADVANCED: Fan-out, Fan-in Pattern
func fanOutFanIn() {
	fmt.Println("\n=== FAN-OUT, FAN-IN PATTERN ===")

	// Input channel
	input := make(chan int)

	// Fan-out: Buat multiple workers
	const numWorkers = 3
	channels := make([]<-chan int, numWorkers)

	for i := 0; i < numWorkers; i++ {
		channels[i] = process(input)
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

func process(input <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for num := range input {
			output <- num * num
		}
	}()
	return output
}

func merge(channels ...<-chan int) <-chan int {
	output := make(chan int)

	go func() {
		defer close(output)
		for _, ch := range channels {
			for val := range ch {
				output <- val
			}
		}
	}()

	return output
}

func main() {
	fmt.Println("=== GO CONCURRENCY EXAMPLES ===")

	// 1. Basic Goroutine
	basicGoroutine()

	// 2. Channel Examples
	channelExamples()

	// 3. Producer-Consumer
	producerConsumer()

	// 4. Buffered vs Unbuffered
	bufferedVsUnbuffered()

	// 5. Select
	selectExample()

	// 6. Worker Pool
	workerPool()

	// 7. Fan-out, Fan-in
	fanOutFanIn()

	fmt.Println("\n=== DONE ===")
}
