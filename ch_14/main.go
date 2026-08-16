package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

// ============================================
// CHAPTER 14: The Context
// ============================================

func main() {
	// =====================================
	// 1. BASIC CONTEXT
	// =====================================
	fmt.Println("=== 1. Basic Context ===")

	ctx := context.Background() // Empty starting context
	fmt.Println("Created background context:", ctx)

	// =====================================
	// 2. CANCEL MANUALLY
	// =====================================
	fmt.Println("\n=== 2. WithCancel (Manual Stop) ===")

	ctx2, cancel := context.WithCancel(context.Background())

	// Worker goroutine
	go func() {
		for i := 1; ; i++ {
			select {
			case <-ctx2.Done(): // Check if cancelled
				fmt.Println("Worker stopped!")
				return
			default:
				fmt.Println("Working...", i)
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	time.Sleep(1 * time.Second)
	cancel() // STOP the worker!
	time.Sleep(100 * time.Millisecond)

	// =====================================
	// 3. TIMEOUT (Auto Cancel)
	// =====================================
	fmt.Println("\n=== 3. WithTimeout (Auto Stop) ===")

	ctx3, cancel3 := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel3()

	// Simulate slow work
	done := make(chan bool)
	go func() {
		time.Sleep(3 * time.Second) // Takes 3 seconds
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("Work finished!")
	case <-ctx3.Done():
		fmt.Println("Timeout! Work took too long")
		fmt.Println("Error:", ctx3.Err())
	}

	// =====================================
	// 4. PASS VALUES
	// =====================================
	fmt.Println("\n=== 4. WithValue (Pass Data) ===")

	ctx4 := context.WithValue(context.Background(), "userID", 123)
	ctx4 = context.WithValue(ctx4, "role", "admin")

	// Get values
	userID := ctx4.Value("userID").(int)
	role := ctx4.Value("role").(string)
	fmt.Println("UserID:", userID)
	fmt.Println("Role:", role)

	// =====================================
	// 5. HTTP REQUEST WITH TIMEOUT
	// =====================================
	fmt.Println("\n=== 5. HTTP with Timeout ===")

	ctx5, cancel5 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel5()

	req, _ := http.NewRequestWithContext(ctx5, "GET", "https://httpbin.org/delay/1", nil)
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("Request failed:", err)
	} else {
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Println("Response length:", len(body), "bytes")
	}

	// =====================================
	// 6. NESTED TIMEOUTS
	// =====================================
	fmt.Println("\n=== 6. Nested Timeouts ===")

	parent, cancelP := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelP()

	child, cancelC := context.WithTimeout(parent, 5*time.Second) // Child wants 5s
	defer cancelC()

	start := time.Now()
	<-child.Done() // Wait for cancel
	fmt.Println("Cancelled after:", time.Since(start).Truncate(time.Second))
	fmt.Println("Child limited by parent (2s, not 5s)")

	// =====================================
	// 7. CHECK WHY CANCELLED
	// =====================================
	fmt.Println("\n=== 7. Check Error Type ===")

	ctx7, cancel7 := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel7()
	time.Sleep(10 * time.Millisecond)

	err = ctx7.Err()
	if err == context.DeadlineExceeded {
		fmt.Println("Cancelled because: Timeout")
	} else if err == context.Canceled {
		fmt.Println("Cancelled because: Manual cancel()")
	}

	// =====================================
	// SUMMARY
	// =====================================
	fmt.Println("\n=== Summary ===")
	fmt.Println("context.Background()     → Starting point")
	fmt.Println("context.WithCancel()     → Manual stop")
	fmt.Println("context.WithTimeout()    → Auto stop after duration")
	fmt.Println("context.WithValue()      → Pass data")
	fmt.Println("ctx.Done()               → Channel closes on cancel")
	fmt.Println("ctx.Err()                → Why cancelled")
}
