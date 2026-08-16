package main

import (
	"fmt"
	"sync"
	"time"
)

// ============================================
// CHAPTER 12: Concurrency in Go
// ============================================

func main() {
	// =====================================
	// 1. GOROUTINE - Run in background
	// =====================================
	fmt.Println("--- 1. Goroutine ---")

	go func() {
		fmt.Println("Hello from goroutine!")
	}()

	fmt.Println("Main continues...")
	time.Sleep(time.Millisecond * 100) // Wait for goroutine

	// =====================================
	// 2. CHANNEL - Send data between goroutines
	// =====================================
	fmt.Println("\n--- 2. Channel ---")

	ch := make(chan string)

	go func() {
		ch <- "Message from goroutine!" // Send
	}()

	msg := <-ch // Receive
	fmt.Println("Received:", msg)

	// =====================================
	// 3. CHANNEL WITH RETURN VALUE
	// =====================================
	fmt.Println("\n--- 3. Channel Return Value ---")

	resultCh := make(chan int)

	go func() {
		result := 10 * 5 // Some work
		resultCh <- result
	}()

	answer := <-resultCh
	fmt.Println("Result:", answer)

	// =====================================
	// 4. BUFFERED CHANNEL
	// =====================================
	fmt.Println("\n--- 4. Buffered Channel ---")

	bufferedCh := make(chan int, 3) // Can hold 3 values

	bufferedCh <- 1
	bufferedCh <- 2
	bufferedCh <- 3
	// bufferedCh <- 4  // Would block!

	fmt.Println(<-bufferedCh) // 1
	fmt.Println(<-bufferedCh) // 2
	fmt.Println(<-bufferedCh) // 3

	// =====================================
	// 5. SELECT - Multiple channels
	// =====================================
	fmt.Println("\n--- 5. Select ---")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch1 <- "from channel 1"
	}()

	go func() {
		time.Sleep(30 * time.Millisecond)
		ch2 <- "from channel 2"
	}()

	// Wait for first response
	select {
	case msg1 := <-ch1:
		fmt.Println("Received:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received:", msg2)
	}

	// =====================================
	// 6. SELECT WITH TIMEOUT
	// =====================================
	fmt.Println("\n--- 6. Select with Timeout ---")

	slowCh := make(chan string)

	go func() {
		time.Sleep(2 * time.Second) // Too slow!
		slowCh <- "finally done"
	}()

	select {
	case msg := <-slowCh:
		fmt.Println("Got:", msg)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Timeout! Too slow.")
	}

	// =====================================
	// 7. MUTEX - Protect shared data
	// =====================================
	fmt.Println("\n--- 7. Mutex ---")

	var counter = 0
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter) // Always 5!

	// =====================================
	// 8. WAITGROUP - Wait for goroutines
	// =====================================
	fmt.Println("\n--- 8. WaitGroup ---")

	var wg2 sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg2.Add(1)
		go func(n int) {
			defer wg2.Done()
			fmt.Printf("Task %d done\n", n)
		}(i)
	}

	wg2.Wait()
	fmt.Println("All tasks completed!")

	// =====================================
	// 9. CLOSE CHANNEL & RANGE
	// =====================================
	fmt.Println("\n--- 9. Close Channel & Range ---")

	numCh := make(chan int)

	go func() {
		for i := 1; i <= 3; i++ {
			numCh <- i
		}
		close(numCh) // Close when done sending
	}()

	for num := range numCh { // Loop until closed
		fmt.Println("Got:", num)
	}

	// =====================================
	// 10. WORKER POOL PATTERN
	// =====================================
	fmt.Println("\n--- 10. Worker Pool ---")

	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// Start 2 workers
	for w := 1; w <= 2; w++ {
		go worker(w, jobs, results)
	}

	// Send 3 jobs
	for j := 1; j <= 3; j++ {
		jobs <- j
	}
	close(jobs)

	// Get results
	for r := 1; r <= 3; r++ {
		fmt.Println("Result:", <-results)
	}

	// =====================================
	// SUMMARY
	// =====================================
	fmt.Println("\n--- Summary ---")
	fmt.Println("go func()        → Goroutine")
	fmt.Println("ch <- value      → Send to channel")
	fmt.Println("<-ch             → Receive from channel")
	fmt.Println("select{}         → Multiple channels")
	fmt.Println("sync.Mutex       → Protect shared data")
	fmt.Println("sync.WaitGroup   → Wait for goroutines")
}

// Worker function for worker pool
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, j)
		results <- j * 2
	}
}
