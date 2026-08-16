package main

// ============================================
// CHAPTER 15: Writing Tests
// ============================================

// Simple functions to test

// Add returns sum of two numbers
func Add(a, b int) int {
	return a + b
}

// Subtract returns difference
func Subtract(a, b int) int {
	return a - b
}

// Multiply returns product
func Multiply(a, b int) int {
	return a * b
}

// Divide returns quotient and error if divide by zero
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

// IsEven checks if number is even
func IsEven(n int) bool {
	return n%2 == 0
}

// Greet returns greeting message
func Greet(name string) string {
	if name == "" {
		return "Hello, stranger!"
	}
	return "Hello, " + name + "!"
}
