package main

import (
	"testing"
)

// ============================================
// CHAPTER 15: Writing Tests
// ============================================

// =====================================
// 1. BASIC TEST
// =====================================

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Error("Expected 5, got", result)
	}
}

// =====================================
// 2. TABLE TEST (Multiple Cases)
// =====================================

func TestAddTable(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{1, 1, 2},
		{2, 3, 5},
		{0, 0, 0},
		{-1, 1, 0},
		{100, 200, 300},
	}

	for _, tc := range tests {
		result := Add(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Add(%d, %d) = %d, want %d",
				tc.a, tc.b, result, tc.expected)
		}
	}
}

// =====================================
// 3. SUBTESTS (Named Cases)
// =====================================

func TestSubtract(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 5, 3, 2},
		{"negative result", 3, 5, -2},
		{"zeros", 0, 0, 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Subtract(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("got %d, want %d", result, tc.expected)
			}
		})
	}
}

// =====================================
// 4. TEST WITH ERROR
// =====================================

func TestDivide(t *testing.T) {
	// Test normal division
	result, err := Divide(10, 2)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	if result != 5 {
		t.Error("Expected 5, got", result)
	}

	// Test divide by zero
	_, err = Divide(10, 0)
	if err == nil {
		t.Error("Expected error for divide by zero")
	}
	if err != ErrDivideByZero {
		t.Error("Expected ErrDivideByZero")
	}
}

// =====================================
// 5. TEST BOOLEAN
// =====================================

func TestIsEven(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{2, true},
		{3, false},
		{0, true},
		{-4, true},
		{-3, false},
	}

	for _, tc := range tests {
		result := IsEven(tc.input)
		if result != tc.expected {
			t.Errorf("IsEven(%d) = %v, want %v",
				tc.input, result, tc.expected)
		}
	}
}

// =====================================
// 6. TEST STRING
// =====================================

func TestGreet(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"with name", "John", "Hello, John!"},
		{"empty name", "", "Hello, stranger!"},
		{"another name", "Jane", "Hello, Jane!"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Greet(tc.input)
			if result != tc.expected {
				t.Errorf("got %q, want %q", result, tc.expected)
			}
		})
	}
}

// =====================================
// 7. BENCHMARK
// =====================================

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(100, 200)
	}
}

func BenchmarkMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiply(100, 200)
	}
}
