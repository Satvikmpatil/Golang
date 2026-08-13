package main

import "fmt"

// ============================================
// CHAPTER 8: Generics
// ============================================

// =====================================
// 1. BASIC GENERIC FUNCTION
// =====================================

// [T any] = T can be any type
func Print[T any](x T) {
	fmt.Println(x)
}

// =====================================
// 2. GENERIC WITH CONSTRAINTS
// =====================================

// Only int or float64 allowed
func Max[T int | float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Add[T int | float64](a, b T) T {
	return a + b
}

// =====================================
// 3. COMPARABLE CONSTRAINT
// =====================================

// comparable = can use == and !=
func Equal[T comparable](a, b T) bool {
	return a == b
}

func Contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// =====================================
// 4. GENERIC STRUCT - Stack
// =====================================

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
	if len(s.items) == 0 {
		var zero T
		return zero
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// =====================================
// 5. GENERIC PAIR
// =====================================

type Pair[T, U any] struct {
	First  T
	Second U
}

func NewPair[T, U any](first T, second U) Pair[T, U] {
	return Pair[T, U]{First: first, Second: second}
}

// =====================================
// 6. FIRST AND LAST ELEMENT
// =====================================

func First[T any](slice []T) T {
	return slice[0]
}

func Last[T any](slice []T) T {
	return slice[len(slice)-1]
}

// =====================================
// 7. MAP FUNCTION (Transform slice)
// =====================================

func Map[T, U any](slice []T, f func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}

// =====================================
// 8. FILTER FUNCTION
// =====================================

func Filter[T any](slice []T, f func(T) bool) []T {
	result := []T{}
	for _, v := range slice {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// =====================================
// MAIN
// =====================================

func main() {
	// --- 1. Basic Generic Function ---
	fmt.Println("--- 1. Print (any type) ---")
	Print(42)
	Print("Hello")
	Print(3.14)
	Print(true)

	// --- 2. Generic with Constraints ---
	fmt.Println("\n--- 2. Max & Add ---")
	fmt.Println("Max(10, 20):", Max(10, 20))
	fmt.Println("Max(3.5, 2.1):", Max(3.5, 2.1))
	fmt.Println("Add(5, 3):", Add(5, 3))
	fmt.Println("Add(1.5, 2.5):", Add(1.5, 2.5))

	// --- 3. Comparable ---
	fmt.Println("\n--- 3. Equal & Contains ---")
	fmt.Println("Equal(10, 10):", Equal(10, 10))
	fmt.Println("Equal(\"a\", \"b\"):", Equal("a", "b"))
	fmt.Println("Contains([1,2,3], 2):", Contains([]int{1, 2, 3}, 2))
	fmt.Println("Contains([1,2,3], 5):", Contains([]int{1, 2, 3}, 5))

	// --- 4. Generic Stack ---
	fmt.Println("\n--- 4. Stack[int] ---")
	intStack := Stack[int]{}
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)
	fmt.Println("Pop:", intStack.Pop())
	fmt.Println("Pop:", intStack.Pop())

	fmt.Println("\n--- Stack[string] ---")
	strStack := Stack[string]{}
	strStack.Push("hello")
	strStack.Push("world")
	fmt.Println("Pop:", strStack.Pop())

	// --- 5. Generic Pair ---
	fmt.Println("\n--- 5. Pair ---")
	p1 := NewPair("name", 25)
	fmt.Println("Pair:", p1.First, p1.Second)

	p2 := NewPair(1, true)
	fmt.Println("Pair:", p2.First, p2.Second)

	// --- 6. First & Last ---
	fmt.Println("\n--- 6. First & Last ---")
	nums := []int{10, 20, 30, 40}
	fmt.Println("First:", First(nums))
	fmt.Println("Last:", Last(nums))

	// --- 7. Map ---
	fmt.Println("\n--- 7. Map (Transform) ---")
	doubled := Map([]int{1, 2, 3}, func(x int) int {
		return x * 2
	})
	fmt.Println("Doubled:", doubled)

	lengths := Map([]string{"go", "rust", "python"}, func(s string) int {
		return len(s)
	})
	fmt.Println("Lengths:", lengths)

	// --- 8. Filter ---
	fmt.Println("\n--- 8. Filter ---")
	evens := Filter([]int{1, 2, 3, 4, 5, 6}, func(x int) bool {
		return x%2 == 0
	})
	fmt.Println("Evens:", evens)

	// --- Summary ---
	fmt.Println("\n--- Summary ---")
	fmt.Println("[T any] = any type")
	fmt.Println("[T comparable] = can use ==")
	fmt.Println("[T int|float64] = specific types")
	fmt.Println("Generics = ONE code, MANY types!")
}
