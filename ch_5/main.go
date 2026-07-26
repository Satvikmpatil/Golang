package main

import "fmt"

// ============================================
// CHAPTER 4: Blocks, Shadows, and Control Structures
// ============================================

// ----- PACKAGE BLOCK (outside any function) -----
var packageVar = 100

func main() {

	// =====================================
	// 1. BLOCKS (Scope)
	// =====================================
	fmt.Println("--- Blocks (Scope) ---")

	// Function block variable
	x := 10
	fmt.Println("x in function:", x)

	if x > 5 {
		// if block - new scope
		y := 20 // Only exists inside if block
		fmt.Println("y in if block:", y)
	}
	// fmt.Println(y)  // ERROR! y doesn't exist here

	// =====================================
	// 2. SHADOWING (Dangerous!)
	// =====================================
	fmt.Println("\n--- Shadowing ---")

	a := 10
	fmt.Println("a before if:", a)

	if a > 5 {
		fmt.Println("a inside if (before shadow):", a) // 10
		a := 5                                         // NEW a (shadows outer!)
		fmt.Println("a inside if (after shadow):", a)  // 5
	}
	fmt.Println("a after if:", a) // 10 (outer a is back!)

	// =====================================
	// 3. IF STATEMENT
	// =====================================
	fmt.Println("\n--- if Statement ---")

	n := 5

	// Basic if/else
	if n == 0 {
		fmt.Println("zero")
	} else if n > 5 {
		fmt.Println("big")
	} else {
		fmt.Println("small")
	}

	// if with variable declaration (scoped to if/else)
	if num := 10; num > 5 {
		fmt.Println("num is big:", num)
	}
	// fmt.Println(num)  // ERROR! num doesn't exist here

	// =====================================
	// 4. FOR LOOP - Complete (C-style)
	// =====================================
	fmt.Println("\n--- for Loop: C-style ---")

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// =====================================
	// 5. FOR LOOP - Condition Only (like while)
	// =====================================
	fmt.Println("\n--- for Loop: Condition Only ---")

	j := 0
	for j < 3 {
		fmt.Println(j)
		j++
	}

	// =====================================
	// 6. FOR LOOP - Infinite
	// =====================================
	fmt.Println("\n--- for Loop: Infinite (with break) ---")

	count := 0
	for {
		fmt.Println("count:", count)
		count++
		if count >= 3 {
			break
		}
	}

	// =====================================
	// 7. FOR-RANGE - Slice
	// =====================================
	fmt.Println("\n--- for-range: Slice ---")

	nums := []int{10, 20, 30}
	for i, v := range nums {
		fmt.Println("index:", i, "value:", v)
	}

	// Only value (ignore index)
	fmt.Println("Only values:")
	for _, v := range nums {
		fmt.Println(v)
	}

	// Only index
	fmt.Println("Only indices:")
	for i := range nums {
		fmt.Println(i)
	}

	// =====================================
	// 8. FOR-RANGE - Map
	// =====================================
	fmt.Println("\n--- for-range: Map ---")

	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range m {
		fmt.Println("key:", k, "value:", v)
	}

	// =====================================
	// 9. FOR-RANGE - String
	// =====================================
	fmt.Println("\n--- for-range: String ---")

	for i, r := range "Hi!" {
		fmt.Println("index:", i, "rune:", string(r))
	}

	// =====================================
	// 10. BREAK AND CONTINUE
	// =====================================
	fmt.Println("\n--- break and continue ---")

	for i := 0; i < 10; i++ {
		if i == 3 {
			continue // Skip 3
		}
		if i == 7 {
			break // Stop at 7
		}
		fmt.Println(i)
	}
	// Output: 0 1 2 4 5 6

	// =====================================
	// 11. SWITCH - Basic
	// =====================================
	fmt.Println("\n--- switch: Basic ---")

	day := 2

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3, 4, 5: // Multiple values
		fmt.Println("Mid-week")
	default:
		fmt.Println("Weekend")
	}

	// =====================================
	// 12. SWITCH - Empty Case
	// =====================================
	fmt.Println("\n--- switch: Empty Case ---")

	size := 7

	switch size {
	case 1, 2, 3:
		fmt.Println("small")
	case 6, 7, 8:
		// Empty! Does nothing
	default:
		fmt.Println("other")
	}
	// Nothing prints for size 7!

	// =====================================
	// 13. SWITCH - Blank (No Expression)
	// =====================================
	fmt.Println("\n--- switch: Blank (like if/else) ---")

	score := 85

	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("C")
	default:
		fmt.Println("F")
	}

	// =====================================
	// 14. SWITCH - With Variable Declaration
	// =====================================
	fmt.Println("\n--- switch: With Variable ---")

	switch num := 15; {
	case num < 10:
		fmt.Println("small")
	case num < 20:
		fmt.Println("medium")
	default:
		fmt.Println("large")
	}

	// =====================================
	// 15. SWITCH - fallthrough (Avoid!)
	// =====================================
	fmt.Println("\n--- switch: fallthrough (avoid!) ---")

	val := 1

	switch val {
	case 1:
		fmt.Println("one")
		fallthrough // Continue to next case
	case 2:
		fmt.Println("two")
	}
	// Output: one, two (both print!)

	// =====================================
	// 16. LABEL - break Out of For Inside Switch
	// =====================================
	fmt.Println("\n--- Label: break loop ---")

	// Problem: break inside switch only breaks switch, not for!
	// Solution: use label

loop:
	for i := 0; i < 10; i++ {
		switch i {
		case 5:
			fmt.Println("Breaking loop at 5")
			break loop // Breaks FOR loop, not just switch!
		default:
			fmt.Println(i)
		}
	}

	// =====================================
	// 17. GOTO (Avoid!)
	// =====================================
	fmt.Println("\n--- goto (avoid!) ---")

	b := 10

	if b > 5 {
		goto done
	}

	fmt.Println("This is skipped")

done:
	fmt.Println("Jumped here with goto")

	// =====================================
	// SUMMARY
	// =====================================
	fmt.Println("\n--- Quick Reference ---")
	fmt.Println("Blocks: { } create new scope")
	fmt.Println("Shadowing: := in inner block hides outer var")
	fmt.Println("for: only loop keyword (4 ways)")
	fmt.Println("switch: no break needed, no fall-through")
	fmt.Println("Labels: for breaking out of nested loops")
	fmt.Println("goto: avoid it!")
}
