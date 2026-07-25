package main

import (
	"errors"
	"fmt"
)

// ============================================
// CHAPTER 5: Functions
// ============================================

// ----- STRUCT FOR SIMULATING NAMED PARAMS -----
type Options struct {
	FName string
	LName string
	Age   int
}

type person struct {
	age  int
	name string
}

func main() {

	// =====================================
	// 1. BASIC FUNCTIONS
	// =====================================
	fmt.Println("--- Basic Functions ---")

	// Regular function
	add := func(a, b int) int {
		return a + b
	}

	// Same type params shortcut: (a, b int) instead of (a int, b int)
	mul := func(a, b int) int {
		return a * b
	}

	fmt.Println("add(2, 3):", add(2, 3))
	fmt.Println("mul(2, 3):", mul(2, 3))

	// =====================================
	// 2. SIMULATING NAMED PARAMETERS (with struct)
	// =====================================
	fmt.Println("\n--- Named Params with Struct ---")

	greet := func(opts Options) {
		fmt.Println(opts.FName, opts.LName, opts.Age)
	}

	greet(Options{
		LName: "Patel",
		Age:   50,
	})

	// =====================================
	// 3. VARIADIC PARAMETERS (...)
	// =====================================
	fmt.Println("\n--- Variadic Parameters ---")

	// Function with variadic param (must be last!)
	addTo := func(base int, vals ...int) []int {
		out := make([]int, 0, len(vals))
		for _, v := range vals {
			out = append(out, base+v)
		}
		return out
	}

	fmt.Println("addTo(3):", addTo(3))
	fmt.Println("addTo(3, 2):", addTo(3, 2))
	fmt.Println("addTo(3, 2, 4, 6):", addTo(3, 2, 4, 6))

	// Pass slice with ... suffix
	nums := []int{4, 3}
	fmt.Println("addTo(3, nums...):", addTo(3, nums...))

	// =====================================
	// 4. MULTIPLE RETURN VALUES
	// =====================================
	fmt.Println("\n--- Multiple Return Values ---")

	divAndRemainder := func(num, denom int) (int, int, error) {
		if denom == 0 {
			return 0, 0, errors.New("cannot divide by zero")
		}
		return num / denom, num % denom, nil
	}

	result, remainder, err := divAndRemainder(10, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 3 =", result, "remainder", remainder)
	}

	// Ignore values with _
	res, _, _ := divAndRemainder(10, 3)
	fmt.Println("Result only:", res)

	// =====================================
	// 5. NAMED RETURN VALUES
	// =====================================
	fmt.Println("\n--- Named Return Values ---")

	namedReturn := func(num, denom int) (result int, remainder int, err error) {
		if denom == 0 {
			err = errors.New("cannot divide by zero")
			return result, remainder, err
		}
		result, remainder = num/denom, num%denom
		return result, remainder, err
	}

	r, rem, e := namedReturn(10, 3)
	fmt.Println("Named return:", r, rem, e)

	// =====================================
	// 6. FUNCTIONS AS VALUES
	// =====================================
	fmt.Println("\n--- Functions as Values ---")

	// Function stored in variable
	var myFunc func(string) int = func(s string) int {
		return len(s)
	}
	fmt.Println("len of 'hello':", myFunc("hello"))

	// =====================================
	// 7. PASSING FUNCTIONS AS PARAMETERS
	// =====================================
	fmt.Println("\n--- Passing Functions as Params ---")

	doMath := func(a, b int, op func(int, int) int) int {
		return op(a, b)
	}

	sub := func(x, y int) int { return x - y }
	addOp := func(x, y int) int { return x + y }

	fmt.Println("doMath(10, 5, sub):", doMath(10, 5, sub))
	fmt.Println("doMath(10, 5, add):", doMath(10, 5, addOp))

	// =====================================
	// 8. RETURNING FUNCTIONS
	// =====================================
	fmt.Println("\n--- Returning Functions ---")

	multiplier := func(x int) func(int) int {
		return func(y int) int {
			return x * y
		}
	}

	double := multiplier(2)
	triple := multiplier(3)

	fmt.Println("double(5):", double(5))
	fmt.Println("triple(5):", triple(5))

	// =====================================
	// 9. ANONYMOUS FUNCTIONS
	// =====================================
	fmt.Println("\n--- Anonymous Functions ---")

	// Assign to variable
	anon := func(x int) int {
		return x * 2
	}
	fmt.Println("anon(5):", anon(5))

	// Call immediately (IIFE)
	func(name string) {
		fmt.Println("Hello", name)
	}("Satvik")

	// Package-level style anonymous functions
	var (
		addVar = func(i, j int) int { return i + j }
		subVar = func(i, j int) int { return i - j }
	)
	fmt.Println("addVar(2, 3):", addVar(2, 3))
	fmt.Println("subVar(2, 3):", subVar(2, 3))

	// =====================================
	// 10. CLOSURES
	// =====================================
	fmt.Println("\n--- Closures ---")

	// Closure can access and modify outer variables
	a := 20
	closure := func() {
		fmt.Println("Inside closure, a =", a)
		a = 30 // Modifies outer variable!
	}
	closure()
	fmt.Println("After closure, a =", a)

	// =====================================
	// 11. SHADOWING (Watch out!)
	// =====================================
	fmt.Println("\n--- Shadowing ---")

	b := 20
	shadowFunc := func() {
		b := 30 // := creates NEW variable (shadow!)
		fmt.Println("Inside (shadow), b =", b)
	}
	shadowFunc()
	fmt.Println("After (original), b =", b) // Still 20!

	// =====================================
	// 12. CALL BY VALUE - Basic Types (Copy)
	// =====================================
	fmt.Println("\n--- Call by Value: Basic Types ---")

	modifyBasic := func(i int, s string, p person) {
		i = 100
		s = "changed"
		p.name = "changed"
	}

	myInt := 10
	myStr := "hello"
	myPerson := person{name: "Alice", age: 25}

	modifyBasic(myInt, myStr, myPerson)
	fmt.Println("int:", myInt)            // 10 - unchanged
	fmt.Println("string:", myStr)         // hello - unchanged
	fmt.Println("person:", myPerson.name) // Alice - unchanged

	// =====================================
	// 13. CALL BY VALUE - Maps (Changes Original!)
	// =====================================
	fmt.Println("\n--- Call by Value: Maps ---")

	modMap := func(m map[int]string) {
		m[1] = "changed"
		m[2] = "new"
		delete(m, 3)
	}

	myMap := map[int]string{1: "first", 3: "third"}
	fmt.Println("Before:", myMap)
	modMap(myMap)
	fmt.Println("After:", myMap) // Changed!

	// =====================================
	// 14. CALL BY VALUE - Slices (Tricky!)
	// =====================================
	fmt.Println("\n--- Call by Value: Slices ---")

	modSlice := func(s []int) {
		s[0] = 99         // Changes original (same array)
		s = append(s, 10) // Creates new array (no room)
		s[0] = 100        // Changes NEW array only
	}

	mySlice := []int{1, 2, 3}
	fmt.Println("Before:", mySlice)
	modSlice(mySlice)
	fmt.Println("After:", mySlice) // [99 2 3] - only first change visible!

	// FIX: Return the slice
	modSliceFixed := func(s []int) []int {
		s[0] = 50
		s = append(s, 10)
		return s
	}

	fixedSlice := []int{1, 2, 3}
	fixedSlice = modSliceFixed(fixedSlice) // Assign back!
	fmt.Println("Fixed:", fixedSlice)      // [50 2 3 10]

	// =====================================
	// SUMMARY
	// =====================================
	fmt.Println("\n--- Quick Reference ---")
	fmt.Println("func name(a, b int) int {}  - Basic function")
	fmt.Println("func f(nums ...int) {}      - Variadic")
	fmt.Println("func f() (int, error) {}    - Multiple returns")
	fmt.Println("var f func(int) int         - Function variable")
	fmt.Println("Closures access outer vars")
	fmt.Println(":= in closure creates shadow")
	fmt.Println("Maps/slices share memory, but append needs return")
}
