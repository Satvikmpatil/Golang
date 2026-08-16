package main

import (
	"fmt"
	"net/http"
)

// ============================================
// CHAPTER 15: Writing Tests
// ============================================
//
// Run tests:     go test
// Verbose:       go test -v
// Coverage:      go test -cover
// Race check:    go test -race
// Benchmark:     go test -bench=.
// Specific test: go test -run TestAdd
//
// ============================================

func main() {
	fmt.Println("=== Chapter 15: Writing Tests ===")
	fmt.Println()

	// Demo the functions
	fmt.Println("Add(2, 3) =", Add(2, 3))
	fmt.Println("Subtract(5, 3) =", Subtract(5, 3))
	fmt.Println("Multiply(4, 5) =", Multiply(4, 5))

	result, err := Divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Divide(10, 2) =", result)
	}

	fmt.Println("IsEven(4) =", IsEven(4))
	fmt.Println("IsEven(5) =", IsEven(5))

	fmt.Println("Greet(\"John\") =", Greet("John"))
	fmt.Println("Greet(\"\") =", Greet(""))

	fmt.Println()
	fmt.Println("=== Test Commands ===")
	fmt.Println("go test          → Run all tests")
	fmt.Println("go test -v       → Verbose output")
	fmt.Println("go test -cover   → Show coverage")
	fmt.Println("go test -race    → Find race conditions")
	fmt.Println("go test -bench=. → Run benchmarks")

	// Start HTTP server for handler testing demo
	fmt.Println()
	fmt.Println("=== HTTP Server ===")
	fmt.Println("Starting server on :8080...")
	fmt.Println("Try: http://localhost:8080/hello")
	fmt.Println("Try: http://localhost:8080/greet?name=John")

	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/greet", GreetHandler)
	http.ListenAndServe(":8080", nil)
}
