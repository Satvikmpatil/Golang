package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// ============================================
// CHAPTER 16: Reflect, Unsafe, and Cgo
// (Advanced - Use with caution!)
// ============================================

func main() {
	// =====================================
	// 1. REFLECT - Inspect types at runtime
	// =====================================
	fmt.Println("=== 1. Reflect ===")

	// Get type and value info
	var x int = 42
	fmt.Println("Type:", reflect.TypeOf(x))   // int
	fmt.Println("Value:", reflect.ValueOf(x)) // 42
	fmt.Println("Kind:", reflect.TypeOf(x).Kind()) // int

	// Inspect struct
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	p := Person{Name: "John", Age: 25}
	inspectStruct(p)

	// =====================================
	// 2. REFLECT - Modify values
	// =====================================
	fmt.Println("\n=== 2. Reflect - Modify ===")

	var num int = 10
	fmt.Println("Before:", num)

	// Must pass pointer to modify
	v := reflect.ValueOf(&num)
	v.Elem().SetInt(100)
	fmt.Println("After:", num)

	// =====================================
	// 3. REFLECT - Create new values
	// =====================================
	fmt.Println("\n=== 3. Reflect - Create ===")

	// Create new int
	newInt := reflect.New(reflect.TypeOf(0))
	newInt.Elem().SetInt(999)
	fmt.Println("New int:", newInt.Elem().Int())

	// =====================================
	// 4. UNSAFE - Get size of types
	// =====================================
	fmt.Println("\n=== 4. Unsafe - Sizeof ===")

	fmt.Println("Size of int:", unsafe.Sizeof(int(0)))
	fmt.Println("Size of int64:", unsafe.Sizeof(int64(0)))
	fmt.Println("Size of string:", unsafe.Sizeof(""))
	fmt.Println("Size of bool:", unsafe.Sizeof(true))

	type Data struct {
		a bool   // 1 byte
		b int64  // 8 bytes
		c bool   // 1 byte
	}
	fmt.Println("Size of Data struct:", unsafe.Sizeof(Data{}))

	// =====================================
	// 5. UNSAFE - Struct field offsets
	// =====================================
	fmt.Println("\n=== 5. Unsafe - Offsetof ===")

	type Example struct {
		A int32
		B int64
		C int32
	}

	fmt.Println("Offset of A:", unsafe.Offsetof(Example{}.A))
	fmt.Println("Offset of B:", unsafe.Offsetof(Example{}.B))
	fmt.Println("Offset of C:", unsafe.Offsetof(Example{}.C))

	// =====================================
	// 6. UNSAFE - Pointer conversion
	// =====================================
	fmt.Println("\n=== 6. Unsafe - Pointers ===")

	var i int64 = 0x0102030405060708
	ptr := unsafe.Pointer(&i)

	// Convert to *byte to read first byte
	b := *(*byte)(ptr)
	fmt.Printf("First byte of %x: %x\n", i, b)

	// =====================================
	// 7. CGO - (Info only, requires C compiler)
	// =====================================
	fmt.Println("\n=== 7. Cgo (Info) ===")
	fmt.Println("Cgo lets you call C code from Go")
	fmt.Println("Requires: C compiler installed")
	fmt.Println("Use case: Integrate with C libraries")
	fmt.Println("Note: Slower than pure Go!")

	// Example cgo code (not runnable without C):
	/*
	// #include <stdio.h>
	// #include <math.h>
	import "C"

	func main() {
	    result := C.sqrt(16)  // Call C's sqrt
	    fmt.Println(result)   // 4
	}
	*/

	// =====================================
	// SUMMARY
	// =====================================
	fmt.Println("\n=== Summary ===")
	fmt.Println("reflect.TypeOf(x)    → Get type")
	fmt.Println("reflect.ValueOf(x)   → Get value")
	fmt.Println("unsafe.Sizeof(x)     → Size in bytes")
	fmt.Println("unsafe.Pointer(x)    → Raw pointer")
	fmt.Println("import \"C\"          → Use C code")
	fmt.Println()
	fmt.Println("⚠️  WARNING: Use these only when absolutely necessary!")
}

// inspectStruct uses reflection to examine a struct
func inspectStruct(s any) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	fmt.Println("\nStruct:", t.Name())
	fmt.Println("Fields:")

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		tag := field.Tag.Get("json")

		fmt.Printf("  - %s: %v (type: %s, tag: %s)\n",
			field.Name, value, field.Type, tag)
	}
}
