# Chapter 3: Composite Types

## Arrays (Fixed Size - Rarely Used!)

```go
var a [3]int                      // [0, 0, 0]
b := [3]int{10, 20, 30}           // with values
c := [...]int{10, 20, 30}         // [...] = Go counts
d := [12]int{1, 5: 4, 6, 10: 100} // sparse array
```

| Feature | Array |
|---------|-------|
| Size | Fixed at compile time |
| `[3]int` vs `[4]int` | Different types! |
| Can compare | ✅ Yes with `==` |
| Use in Go | ❌ Rarely |

**Why rarely used?** Size is part of type. Use slices instead!

---

## Slices (Flexible - Use This!)

```go
var s []int                // nil slice
s := []int{1, 2, 3}        // slice literal
s := make([]int, 5)        // length=5 (all zeros)
s := make([]int, 0, 10)    // length=0, capacity=10
```

| `[]` | `[3]` or `[...]` |
|------|------------------|
| Slice | Array |
| Flexible | Fixed |

---

## len and cap

```go
s := []int{1, 2, 3}
len(s)    // 3 - how many items
cap(s)    // 3 - space reserved
```

---

## append

```go
x := []int{1, 2}
x = append(x, 3)           // [1, 2, 3]
x = append(x, 4, 5, 6)     // add multiple
x = append(x, y...)        // append slice (use ...)
```

**MUST assign back!** `x = append(x, 3)`

### Common Mistake

```go
// WRONG
x := make([]int, 5)        // [0,0,0,0,0]
x = append(x, 10)          // [0,0,0,0,0,10] - adds after zeros!

// CORRECT
x := make([]int, 0, 5)     // length=0
x = append(x, 10)          // [10]
```

---

## Slicing

```go
x := []string{"a", "b", "c", "d", "e"}

x[:2]     // [a, b]         from start to index 1
x[2:]     // [c, d, e]      from index 2 to end
x[1:3]    // [b, c]         from index 1 to 2
x[:]      // [a, b, c, d, e] everything
```

---

## Slices Share Memory! (Important!)

```go
x := []string{"a", "b", "c", "d"}
y := x[:2]       // y = [a, b]

y[0] = "X"       // change y
fmt.Println(x)   // [X, b, c, d] ← x changed too!
```

### Safe Way: Full Slice Expression

```go
y := x[:2:2]     // length=2, capacity=2 (limited)
y = append(y, "z")
// x unchanged because y got new memory
```

---

## clear (Empty a Slice)

```go
s := []string{"a", "b", "c"}
clear(s)
// s = ["", "", ""] - zeros, but length still 3
```

---

## Compare Slices

```go
import "slices"

a := []int{1, 2, 3}
b := []int{1, 2, 3}
slices.Equal(a, b)    // true

// CAN'T use a == b with slices!
```

---

## Array ↔ Slice

| Conversion | Memory |
|------------|--------|
| Array → Slice `arr[:]` | SHARED |
| Slice → Array `[4]int(slice)` | COPIED |

```go
// Array to Slice (shares memory!)
arr := [3]int{1, 2, 3}
slc := arr[:]
slc[0] = 100    // arr also changes!

// Slice to Array (copies!)
slc := []int{1, 2, 3}
arr := [3]int(slc)
slc[0] = 100    // arr NOT affected
```

---

## Strings, Bytes, Runes

**String = sequence of bytes (not characters!)**

```go
s := "Hello"
len(s)           // 5 bytes

s := "Hello ☀"
len(s)           // 9 bytes (emoji = 4 bytes!)
```

### Safe Character Count

```go
import "unicode/utf8"
utf8.RuneCountInString("Hello ☀")  // 7 characters
```

### Conversions

```go
// String ↔ Bytes
bytes := []byte("Hello")
str := string(bytes)

// String ↔ Runes (safe for Unicode)
runes := []rune("Hello ☀")
len(runes)  // 7 (correct character count!)
```

---

## Maps (Key-Value Pairs)

```go
// Create
m := map[string]int{}
m := map[string]int{"a": 1, "b": 2}

// Read & Write
m["key"] = 10
v := m["key"]

// Delete
delete(m, "key")

// Clear
clear(m)
```

### Zero Value = nil

```go
var m map[string]int   // nil
m["key"] = 1           // PANIC!

m = map[string]int{}   // initialize first
m["key"] = 1           // OK
```

---

## Comma OK Idiom

```go
m := map[string]int{"a": 1, "b": 0}

v, ok := m["a"]    // v=1, ok=true
v, ok := m["b"]    // v=0, ok=true (exists with value 0!)
v, ok := m["c"]    // v=0, ok=false (doesn't exist)

// Common pattern
if v, ok := m["key"]; ok {
    fmt.Println("Found:", v)
}
```

---

## Compare Maps

```go
import "maps"

m1 := map[string]int{"a": 1}
m2 := map[string]int{"a": 1}
maps.Equal(m1, m2)    // true
```

---

## Map as Set

```go
set := map[int]bool{}
set[5] = true
set[10] = true

if set[5] {
    fmt.Println("5 exists")
}
```

---

## Structs

```go
// Define
type person struct {
    name string
    age  int
}

// Create
var p person                      // zero value
p := person{}                     // same as above
p := person{"Bob", 25}            // with values (order matters)
p := person{name: "Bob", age: 25} // with names (recommended)

// Read & Write
p.name = "Alice"
fmt.Println(p.name)
```

---

## Anonymous Struct

```go
pet := struct {
    name string
    kind string
}{
    name: "Fido",
    kind: "dog",
}
```

---

## Comparing Structs

```go
p1 := point{1, 2}
p2 := point{1, 2}
p1 == p2    // true (if all fields comparable)
```

---

## Key Takeaways

1. **Use slices, not arrays** - slices are flexible
2. **`[]` = slice, `[3]` = array** - remember the difference
3. **append returns new slice** - must assign back
4. **Slices share memory** - be careful when slicing
5. **Full slice `x[:2:2]`** - limits capacity, prevents sharing
6. **String length = bytes** - use `utf8.RuneCountInString` for chars
7. **Map zero value = nil** - initialize before using
8. **Comma ok idiom** - check if map key exists
9. **Struct groups data** - like a class without methods

---

## Quick Reference

| Type | Zero Value | Compare |
|------|------------|---------|
| Array `[3]int` | All zeros | ✅ `==` |
| Slice `[]int` | `nil` | ❌ Use `slices.Equal` |
| Map `map[K]V` | `nil` | ❌ Use `maps.Equal` |
| Struct | All fields zero | ✅ `==` (if fields comparable) |

| Create | Syntax |
|--------|--------|
| Array | `[3]int{1,2,3}` |
| Slice | `[]int{1,2,3}` or `make([]int, len, cap)` |
| Map | `map[string]int{}` or `make(map[string]int)` |
| Struct | `person{name: "Bob"}` |
