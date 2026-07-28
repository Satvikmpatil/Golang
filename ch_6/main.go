package main

import "fmt"

// ============================================
// CHAPTER 6: Pointers
// ============================================

func main() {

	// =====================================
	// 1. POINTER BASICS: & and *
	// =====================================
	fmt.Println("--- 1. Pointer Basics ---")

	x := 10
	ptr := &x // & = get address

	fmt.Println("x =", x)
	fmt.Println("ptr (address) =", ptr)
	fmt.Println("*ptr (value) =", *ptr)

	// Change value through pointer
	*ptr = 50
	fmt.Println("After *ptr = 50, x =", x) // 50 (changed!)

	// =====================================
	// 2. POINTER TYPE DECLARATION
	// =====================================
	fmt.Println("\n--- 2. Pointer Type ---")

	var ptr2 *int // *int = pointer to int
	ptr2 = &x
	fmt.Println("ptr2 =", ptr2)
	fmt.Println("*ptr2 =", *ptr2)

	// =====================================
	// 3. new() FUNCTION
	// =====================================
	fmt.Println("\n--- 3. new() Function ---")

	ptr3 := new(int) // Creates pointer to zero-value
	fmt.Println("ptr3 =", ptr3)
	fmt.Println("*ptr3 =", *ptr3) // 0 (zero value)

	// =====================================
	// 4. NIL POINTER
	// =====================================
	fmt.Println("\n--- 4. nil Pointer ---")

	var ptr4 *int // nil by default
	fmt.Println("ptr4 == nil?", ptr4 == nil)

	// Safe way to use
	if ptr4 != nil {
		fmt.Println(*ptr4)
	} else {
		fmt.Println("ptr4 is nil, can't dereference!")
	}

	// =====================================
	// 5. CAN'T ADDRESS LITERALS
	// =====================================
	fmt.Println("\n--- 5. Can't Address Literals ---")

	// ptr := &10  // Error! Can't do this

	// Solution: Helper function
	makePointer := func(t int) *int {
		return &t
	}

	ptr5 := makePointer(10)
	fmt.Println("*ptr5 =", *ptr5)

	// =====================================
	// 6. STRUCT WITH POINTER FIELD
	// =====================================
	fmt.Println("\n--- 6. Struct with Pointer Field ---")

	type Person struct {
		Name string
		Age  *int
	}

	// Method 1: Use variable
	age := 25
	p1 := Person{Name: "Satvik", Age: &age}
	fmt.Println("p1:", p1.Name, *p1.Age)

	// Method 2: Use helper function
	p2 := Person{Name: "Souju", Age: makePointer(30)}
	fmt.Println("p2:", p2.Name, *p2.Age)

	// =====================================
	// 7. FUNCTION WITHOUT POINTER (NO CHANGE)
	// =====================================
	fmt.Println("\n--- 7. Function Without Pointer ---")

	changeValue := func(n int) {
		n = 100 // Changes copy only!
	}

	r := 10
	changeValue(r)
	fmt.Println("After changeValue(r):", r) // Still 10!

	// =====================================
	// 8. FUNCTION WITH POINTER (CHANGES!)
	// =====================================
	fmt.Println("\n--- 8. Function With Pointer ---")

	changePointer := func(n *int) {
		*n = 100 // Changes original!
	}

	t := 10
	changePointer(&t)
	fmt.Println("After changePointer(&t):", t) // 100!

	// =====================================
	// 9. MAP IN FUNCTION (ALWAYS CHANGES)
	// =====================================
	fmt.Println("\n--- 9. Map in Function ---")

	modifyMap := func(m map[string]int) {
		m["a"] = 100
		m["b"] = 200
	}

	myMap := map[string]int{"a": 1}
	fmt.Println("Before:", myMap)
	modifyMap(myMap)
	fmt.Println("After:", myMap) // Changed!

	// =====================================
	// 10. SLICE IN FUNCTION (TRICKY!)
	// =====================================
	fmt.Println("\n--- 10. Slice in Function ---")

	modifySlice := func(s []int) {
		s[0] = 100       // Changes original!
		s = append(s, 5) // Doesn't change original!
	}

	mySlice := []int{1, 2, 3}
	fmt.Println("Before:", mySlice)
	modifySlice(mySlice)
	fmt.Println("After:", mySlice) // [100 2 3]

	// =====================================
	// 11. WRONG WAY: REASSIGN POINTER
	// =====================================
	fmt.Println("\n--- 11. Wrong Way (Reassign Pointer) ---")

	failedUpdate := func(px *int) {
		x2 := 20
		px = &x2 // Only changes copy of pointer!
	}

	x1 := 10
	failedUpdate(&x1)
	fmt.Println("After failedUpdate:", x1) // Still 10!

	// =====================================
	// 12. RIGHT WAY: CHANGE VALUE AT POINTER
	// =====================================
	fmt.Println("\n--- 12. Right Way (Change Value) ---")

	update := func(px *int) {
		*px = 20 // Changes actual value!
	}

	x2 := 10
	update(&x2)
	fmt.Println("After update:", x2) // 20!

	// =====================================
	// 13. BAD: CREATING GARBAGE IN LOOP
	// =====================================
	fmt.Println("\n--- 13. Bad: Garbage in Loop ---")

	bad := func() {
		for i := 0; i < 3; i++ {
			p := &Person{Name: "test"} // New allocation every time!
			fmt.Println("bad:", p.Name)
		}
	}
	bad()

	// =====================================
	// 14. GOOD: REUSE STRUCT
	// =====================================
	fmt.Println("\n--- 14. Good: Reuse Struct ---")

	good := func() {
		p := Person{} // 1 allocation only
		for i := 0; i < 3; i++ {
			p.Name = "test" // Reuse same struct
			fmt.Println("good:", p.Name)
		}
	}
	good()

	// =====================================
	// SUMMARY
	// =====================================
	fmt.Println("\n--- Summary ---")
	fmt.Println("& = get address")
	fmt.Println("* = get value at address")
	fmt.Println("*ptr = 20  → changes original")
	fmt.Println("ptr = &x   → doesn't change original")
	fmt.Println("Map → always changes in function")
	fmt.Println("Slice[i] → changes, append → doesn't")
}
