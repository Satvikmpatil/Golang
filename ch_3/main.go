package main

import (
	"fmt"
	"maps"
	"slices"
	"unicode/utf8"
)

// ============================================
// CHAPTER 3: Composite Types
// ============================================

// ----- STRUCT DEFINITION (at package level) -----
type person struct {
	name string
	age  int
	pet  string
}

type point struct {
	x, y int
}

func main() {

	// =====================================
	// 1. ARRAYS (fixed size, rarely used)
	// =====================================
	fmt.Println("--- Arrays ---")

	var a [3]int                        // [0, 0, 0] - zero values
	b := [3]int{10, 20, 30}             // with values
	c := [...]int{10, 20, 30}           // [...] = Go counts size
	d := [12]int{1, 5: 4, 6, 10: 100}   // sparse array

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)
	fmt.Println("b == c:", b == c)      // true - can compare arrays
	fmt.Println("len(b):", len(b))

	// =====================================
	// 2. SLICES (flexible, use this!)
	// =====================================
	fmt.Println("\n--- Slices ---")

	// Different ways to create slices
	var s1 []int                        // nil slice
	s2 := []int{1, 2, 3}                // slice literal
	s3 := make([]int, 5)                // length=5, capacity=5, all zeros
	s4 := make([]int, 0, 10)            // length=0, capacity=10

	fmt.Println("s1 (nil):", s1, "len:", len(s1), "cap:", cap(s1))
	fmt.Println("s2:", s2, "len:", len(s2), "cap:", cap(s2))
	fmt.Println("s3:", s3, "len:", len(s3), "cap:", cap(s3))
	fmt.Println("s4:", s4, "len:", len(s4), "cap:", cap(s4))
	fmt.Println("s1 == nil:", s1 == nil)

	// =====================================
	// 3. APPEND (add items to slice)
	// =====================================
	fmt.Println("\n--- Append ---")

	x := []int{1, 2}
	x = append(x, 3)           // add one
	x = append(x, 4, 5, 6)     // add multiple
	fmt.Println("after append:", x)

	y := []int{7, 8, 9}
	x = append(x, y...)        // append another slice (use ...)
	fmt.Println("after append slice:", x)

	// MISTAKE: append on make([]int, 5) adds AFTER zeros!
	mistake := make([]int, 5)
	mistake = append(mistake, 10)
	fmt.Println("mistake:", mistake) // [0 0 0 0 0 10]

	// CORRECT: use length=0
	correct := make([]int, 0, 5)
	correct = append(correct, 10)
	fmt.Println("correct:", correct) // [10]

	// =====================================
	// 4. SLICING (get part of slice/array)
	// =====================================
	fmt.Println("\n--- Slicing ---")

	letters := []string{"a", "b", "c", "d", "e"}
	fmt.Println("original:", letters)
	fmt.Println("[:2]  :", letters[:2])   // [a b]
	fmt.Println("[2:]  :", letters[2:])   // [c d e]
	fmt.Println("[1:3] :", letters[1:3])  // [b c]
	fmt.Println("[:]   :", letters[:])    // [a b c d e]

	// =====================================
	// 5. SLICES SHARE MEMORY (careful!)
	// =====================================
	fmt.Println("\n--- Slices Share Memory ---")

	original := []string{"a", "b", "c", "d"}
	sub := original[:2]        // sub shares memory with original!

	sub[0] = "X"               // change sub
	fmt.Println("original:", original) // [X b c d] - changed!
	fmt.Println("sub:", sub)           // [X b]

	// SAFE WAY: full slice expression limits capacity
	safe := []string{"a", "b", "c", "d"}
	safeSub := safe[:2:2]      // length=2, capacity=2 (limited!)
	safeSub = append(safeSub, "z")
	fmt.Println("safe:", safe)         // [a b c d] - unchanged!
	fmt.Println("safeSub:", safeSub)   // [a b z]

	// =====================================
	// 6. CLEAR (empty a slice)
	// =====================================
	fmt.Println("\n--- Clear ---")

	toClear := []string{"a", "b", "c"}
	fmt.Println("before clear:", toClear, "len:", len(toClear))
	clear(toClear)
	fmt.Println("after clear:", toClear, "len:", len(toClear)) // length same!

	// =====================================
	// 7. COMPARE SLICES
	// =====================================
	fmt.Println("\n--- Compare Slices ---")

	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}
	slice3 := []int{1, 2, 4}

	// Can't use == with slices! Use slices.Equal
	fmt.Println("slice1 == slice2:", slices.Equal(slice1, slice2)) // true
	fmt.Println("slice1 == slice3:", slices.Equal(slice1, slice3)) // false

	// =====================================
	// 8. ARRAY ↔ SLICE CONVERSION
	// =====================================
	fmt.Println("\n--- Array ↔ Slice ---")

	// Array to Slice (SHARES memory)
	arr := [4]int{1, 2, 3, 4}
	sliceFromArr := arr[:]
	sliceFromArr[0] = 100
	fmt.Println("arr:", arr)                   // [100 2 3 4] - changed!
	fmt.Println("sliceFromArr:", sliceFromArr)

	// Slice to Array (COPIES memory)
	slc := []int{5, 6, 7, 8}
	arrFromSlice := [4]int(slc)
	slc[0] = 500
	fmt.Println("slc:", slc)                   // [500 6 7 8]
	fmt.Println("arrFromSlice:", arrFromSlice) // [5 6 7 8] - not changed!

	// =====================================
	// 9. STRINGS, BYTES, RUNES
	// =====================================
	fmt.Println("\n--- Strings, Bytes, Runes ---")

	str := "Hello"
	fmt.Println("str:", str)
	fmt.Println("len(str):", len(str))     // 5 bytes
	fmt.Println("str[0]:", str[0])         // 72 (byte value of 'H')
	fmt.Println("str[1:4]:", str[1:4])     // "ell"

	// Emoji = multiple bytes!
	emoji := "Hello ☀"
	fmt.Println("emoji:", emoji)
	fmt.Println("len(emoji):", len(emoji)) // 10 bytes (not 7 chars!)
	fmt.Println("rune count:", utf8.RuneCountInString(emoji)) // 7 chars

	// String to bytes
	bytes := []byte("Hello")
	fmt.Println("bytes:", bytes)           // [72 101 108 108 111]

	// String to runes (safe for Unicode)
	runes := []rune("Hello ☀")
	fmt.Println("runes:", runes)           // [72 101 108 108 111 32 9728]
	fmt.Println("len(runes):", len(runes)) // 7 (correct!)

	// =====================================
	// 10. MAPS (key-value pairs)
	// =====================================
	fmt.Println("\n--- Maps ---")

	// Create and use map
	scores := map[string]int{}
	scores["Alice"] = 100
	scores["Bob"] = 85
	fmt.Println("scores:", scores)
	fmt.Println("Alice:", scores["Alice"])
	fmt.Println("Unknown:", scores["Unknown"]) // 0 (zero value)

	// Map with initial values
	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}
	fmt.Println("ages:", ages)

	// =====================================
	// 11. COMMA OK IDIOM (check if key exists)
	// =====================================
	fmt.Println("\n--- Comma OK ---")

	m := map[string]int{
		"exists":    5,
		"alsoZero":  0,
	}

	v1, ok1 := m["exists"]    // v1=5, ok1=true
	v2, ok2 := m["alsoZero"]  // v2=0, ok2=true (exists!)
	v3, ok3 := m["missing"]   // v3=0, ok3=false (doesn't exist)

	fmt.Println("exists:", v1, ok1)
	fmt.Println("alsoZero:", v2, ok2)
	fmt.Println("missing:", v3, ok3)

	// Common pattern
	if val, ok := m["exists"]; ok {
		fmt.Println("Found:", val)
	}

	// =====================================
	// 12. DELETE AND CLEAR MAPS
	// =====================================
	fmt.Println("\n--- Delete & Clear Maps ---")

	toDelete := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("before delete:", toDelete)
	delete(toDelete, "b")
	fmt.Println("after delete:", toDelete)
	clear(toDelete)
	fmt.Println("after clear:", toDelete, "len:", len(toDelete))

	// =====================================
	// 13. COMPARE MAPS
	// =====================================
	fmt.Println("\n--- Compare Maps ---")

	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map[string]int{"b": 2, "a": 1}
	map3 := map[string]int{"a": 1, "b": 3}

	fmt.Println("map1 == map2:", maps.Equal(map1, map2)) // true
	fmt.Println("map1 == map3:", maps.Equal(map1, map3)) // false

	// =====================================
	// 14. MAP AS SET (no duplicates)
	// =====================================
	fmt.Println("\n--- Map as Set ---")

	set := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}

	for _, v := range vals {
		set[v] = true
	}

	fmt.Println("vals length:", len(vals))  // 11
	fmt.Println("set length:", len(set))    // 8 (no duplicates!)

	if set[5] {
		fmt.Println("5 is in set")
	}

	// =====================================
	// 15. STRUCTS
	// =====================================
	fmt.Println("\n--- Structs ---")

	// Zero value struct
	var fred person
	fmt.Println("fred (zero):", fred)

	// Empty literal (same as zero)
	bob := person{}
	fmt.Println("bob (empty):", bob)

	// With values (order matters)
	julia := person{"Julia", 40, "cat"}
	fmt.Println("julia:", julia)

	// With field names (recommended!)
	beth := person{
		name: "Beth",
		age:  30,
	}
	fmt.Println("beth:", beth)

	// Read and write fields
	bob.name = "Bob"
	bob.age = 25
	bob.pet = "dog"
	fmt.Println("bob updated:", bob)
	fmt.Println("bob's name:", bob.name)

	// =====================================
	// 16. ANONYMOUS STRUCT
	// =====================================
	fmt.Println("\n--- Anonymous Struct ---")

	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}
	fmt.Println("pet:", pet)
	fmt.Println("pet name:", pet.name)

	// =====================================
	// 17. COMPARING STRUCTS
	// =====================================
	fmt.Println("\n--- Comparing Structs ---")

	p1 := point{1, 2}
	p2 := point{1, 2}
	p3 := point{3, 4}

	fmt.Println("p1 == p2:", p1 == p2) // true
	fmt.Println("p1 == p3:", p1 == p3) // false

	// =====================================
	// SUMMARY
	// =====================================
	fmt.Println("\n--- Quick Reference ---")
	fmt.Println("Array:  [3]int{1,2,3}    - fixed size")
	fmt.Println("Slice:  []int{1,2,3}     - flexible (use this!)")
	fmt.Println("Map:    map[string]int{} - key-value pairs")
	fmt.Println("Struct: type X struct{}  - group related data")
}
