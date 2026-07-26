# Chapter 4: Blocks, Shadows, and Control Structures

## Blocks (Scope)

**Block = Area where variables exist `{ }`**

```go
package main           // Package block

var x = 10             // Package block variable

func main() {          // Function block
    y := 20            // Function block variable
    
    if y > 10 {        // if block
        z := 30        // Only exists inside if
    }
    // z doesn't exist here!
}
```

**Rule:** Inner can see outer, outer can't see inner.

---

## Shadowing (Dangerous!)

**Same name in inner block "hides" outer variable**

```go
a := 10
if a > 5 {
    fmt.Println(a)   // 10 (outer)
    a := 5           // NEW a (shadows outer!)
    fmt.Println(a)   // 5 (inner)
}
fmt.Println(a)       // 10 (outer is back!)
```

### `:=` Trap

```go
x := 10
if true {
    x, y := 5, 20    // x is shadowed! (y is new)
}
fmt.Println(x)       // Still 10!
```

### Don't Shadow These!

- `true`, `false`, `nil`
- `int`, `string`, `bool`
- `fmt`, `errors`
- `make`, `len`, `cap`

---

## if Statement

```go
// Basic
if n == 0 {
    fmt.Println("zero")
} else if n > 5 {
    fmt.Println("big")
} else {
    fmt.Println("small")
}

// With variable declaration (scoped to if/else)
if num := 10; num > 5 {
    fmt.Println(num)
}
// num doesn't exist here!
```

**No parentheses around condition!**

---

## for Loop - 4 Ways

Go has only `for`, no `while`!

### 1. Complete (C-style)

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

### 2. Condition Only (like while)

```go
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}
```

### 3. Infinite Loop

```go
for {
    fmt.Println("forever")
    if condition {
        break
    }
}
```

### 4. for-range

```go
// Slice
nums := []int{10, 20, 30}
for i, v := range nums {
    fmt.Println(i, v)   // index, value
}

// Map
m := map[string]int{"a": 1}
for k, v := range m {
    fmt.Println(k, v)   // key, value
}

// String
for i, r := range "hello" {
    fmt.Println(i, string(r))  // index, rune
}
```

### Ignore with `_`

```go
// Only value
for _, v := range nums { }

// Only index
for i := range nums { }
```

---

## break and continue

```go
for i := 0; i < 10; i++ {
    if i == 3 {
        continue    // Skip this iteration
    }
    if i == 7 {
        break       // Exit loop
    }
    fmt.Println(i)
}
// Output: 0 1 2 4 5 6
```

---

## switch Statement

### Basic

```go
switch n {
case 1:
    fmt.Println("one")
case 2, 3, 4:              // Multiple values
    fmt.Println("two to four")
default:
    fmt.Println("other")
}
```

**No `break` needed!** Go doesn't fall through.

### Empty Case = Do Nothing

```go
switch size {
case 6, 7, 8:
    // Empty! Nothing happens
default:
    fmt.Println("other")
}
```

### Blank Switch (No Expression)

```go
switch {
case n < 10:
    fmt.Println("small")
case n < 20:
    fmt.Println("medium")
default:
    fmt.Println("large")
}
```

### With Variable Declaration

```go
switch num := 15; {
case num < 10:
    fmt.Println("small")
default:
    fmt.Println("big")
}
```

### fallthrough (Avoid!)

```go
switch val {
case 1:
    fmt.Println("one")
    fallthrough          // Continue to next case
case 2:
    fmt.Println("two")
}
// val=1 prints: one, two (both!)
```

**Don't use it!** Restructure your code instead.

---

## Labels

### Problem: break in Switch Inside For

```go
for i := 0; i < 10; i++ {
    switch i {
    case 7:
        break    // Only breaks SWITCH, not FOR!
    }
}
// Loop continues!
```

### Solution: Use Label

```go
loop:
for i := 0; i < 10; i++ {
    switch i {
    case 7:
        break loop   // Breaks FOR loop!
    }
}
```

---

## goto (Avoid!)

```go
if condition {
    goto done
}
fmt.Println("skipped")

done:
fmt.Println("jumped here")
```

**Rules:**
- Can't skip variable declarations
- Can't jump into blocks

**Almost never use it!**

---

## Quick Summary

| Concept | Remember |
|---------|----------|
| Block | `{ }` creates new scope |
| Shadowing | `:=` in inner block creates new variable |
| if | No parentheses, can declare variable |
| for | Only loop (4 ways) |
| for-range | For slice, map, string |
| switch | No break needed, no fall-through |
| Empty case | Does nothing |
| Blank switch | `switch { case n < 5: }` |
| Labels | For breaking nested loops |
| goto | Avoid it! |

---

## Key Takeaways

1. **Blocks define scope** - variables exist only in their block

2. **`:=` can shadow** - be careful in inner blocks

3. **Never shadow** `true`, `false`, `nil`, `fmt`, etc.

4. **Only `for`** - no while, no do-while

5. **for-range** - use for collections

6. **switch no break** - doesn't fall through

7. **Empty case** - does nothing

8. **Blank switch** - like if/else but cleaner

9. **Labels** - for breaking out of nested loops

10. **goto** - exists but avoid it!
