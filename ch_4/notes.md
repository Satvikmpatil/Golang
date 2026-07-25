# Chapter 5: Functions

## Basic Function Structure

```go
func functionName(param1 type1, param2 type2) returnType {
    return value
}
```

**Example:**
```go
func add(a int, b int) int {
    return a + b
}

// Shortcut: same type params
func add(a, b int) int {
    return a + b
}
```

---

## No Named/Optional Parameters

Go doesn't have named or optional parameters. Use a struct instead:

```go
type Options struct {
    FirstName string
    LastName  string
    Age       int
}

func greet(opts Options) {
    fmt.Println(opts.FirstName, opts.LastName)
}

// Call - looks like named params!
greet(Options{
    LastName: "Patel",
    Age:      50,
})
```

---

## Variadic Parameters (Any Number of Inputs)

Use `...` before the type:

```go
func addTo(base int, vals ...int) []int {
    out := make([]int, 0, len(vals))
    for _, v := range vals {
        out = append(out, base+v)
    }
    return out
}

// Call it
addTo(3)              // []
addTo(3, 2)           // [5]
addTo(3, 2, 4, 6)     // [5, 7, 9]

// Pass slice - use ... after it!
nums := []int{1, 2, 3}
addTo(3, nums...)     // [4, 5, 6]
```

**Rules:**
- Must be last parameter
- Inside function, it's a slice
- Pass slice with `...` suffix

---

## Multiple Return Values

```go
func divAndRemainder(num, denom int) (int, int, error) {
    if denom == 0 {
        return 0, 0, errors.New("cannot divide by zero")
    }
    return num / denom, num % denom, nil
}

// Call - must assign all values
result, remainder, err := divAndRemainder(5, 2)

// Ignore with _
result, _, _ := divAndRemainder(5, 2)
```

**Convention:** Error is always the LAST return value.

---

## Named Return Values

```go
func divAndRemainder(num, denom int) (result int, remainder int, err error) {
    if denom == 0 {
        err = errors.New("cannot divide by zero")
        return result, remainder, err
    }
    result, remainder = num/denom, num%denom
    return result, remainder, err
}
```

**Note:** Named returns are pre-declared with zero values.

---

## Blank Returns - NEVER USE!

```go
// BAD - Don't do this!
func bad(x int) (result int) {
    result = x * 2
    return    // ← No values specified, confusing!
}
```

Always specify what you're returning for clarity.

---

## Functions Are Values

Functions can be stored in variables:

```go
// Declare function variable
var myFunc func(string) int

// Assign a function
myFunc = func(s string) int {
    return len(s)
}

// Call it
fmt.Println(myFunc("hello"))  // 5
```

**Zero value:** `nil` (will panic if called!)

---

## Passing Functions as Parameters

```go
func doMath(a, b int, operation func(int, int) int) int {
    return operation(a, b)
}

// Use it
add := func(x, y int) int { return x + y }
sub := func(x, y int) int { return x - y }

fmt.Println(doMath(10, 5, add))  // 15
fmt.Println(doMath(10, 5, sub))  // 5
```

---

## Returning Functions

```go
func multiplier(x int) func(int) int {
    return func(y int) int {
        return x * y
    }
}

// Use it
double := multiplier(2)
triple := multiplier(3)

fmt.Println(double(5))   // 10
fmt.Println(triple(5))   // 15
```

---

## Anonymous Functions

Function without a name:

```go
// Assign to variable
f := func(x int) int {
    return x * 2
}

// Call immediately (IIFE)
func(name string) {
    fmt.Println("Hello", name)
}("Satvik")
```

---

## Closures

Inner function that "remembers" outer variables:

```go
func main() {
    a := 20

    f := func() {
        fmt.Println(a)   // Can read outer variable
        a = 30           // Can modify outer variable!
    }

    f()
    fmt.Println(a)       // 30 - changed!
}
```

### Watch Out: Shadowing

```go
a := 20

f := func() {
    a := 30    // ← := creates NEW variable (shadow)!
    fmt.Println(a)
}

f()
fmt.Println(a)   // 20 - unchanged!
```

| `a = 30` | `a := 30` |
|----------|-----------|
| Modifies outer | Creates new (shadow) |

---

## Go Is Call by Value

Go makes a COPY when passing to functions.

### Basic Types - Copy (Original Safe)

```go
func modify(i int, s string) {
    i = 100
    s = "changed"
}

x := 10
str := "hello"
modify(x, str)
fmt.Println(x, str)   // 10 hello - unchanged!
```

### Structs - Copy (Original Safe)

```go
type person struct {
    name string
}

func modify(p person) {
    p.name = "Bob"
}

p := person{name: "Alice"}
modify(p)
fmt.Println(p.name)   // Alice - unchanged!
```

### Maps - Changes Original!

```go
func modMap(m map[int]string) {
    m[1] = "changed"
}

m := map[int]string{1: "first"}
modMap(m)
fmt.Println(m)   // map[1:changed] - CHANGED!
```

### Slices - Tricky!

```go
func modSlice(s []int) {
    s[0] = 99           // ✅ Changes original
    s = append(s, 10)   // ❌ Doesn't change original
    s[0] = 100          // ❌ Changes NEW array only
}

s := []int{1, 2, 3}
modSlice(s)
fmt.Println(s)   // [99 2 3]
```

**Why append doesn't work?**
- Slice = pointer + length + capacity
- Function gets COPY of these 3 values
- `append` may create new array, updates pointer in COPY only
- Original still points to old array

**Fix:** Return the new slice!
```go
func modSlice(s []int) []int {
    s = append(s, 10)
    return s
}

s = modSlice(s)   // Assign back!
```

---

## Quick Reference

| What | Modify in Function | Affects Original? |
|------|-------------------|-------------------|
| `int`, `string`, `bool` | Copy | ❌ No |
| `struct` | Copy | ❌ No |
| `map` elements | Shared | ✅ Yes |
| `slice` elements | Shared | ✅ Yes |
| `slice` append | New array | ❌ No |

---

## Key Takeaways

1. **Function syntax:** `func name(params) returnType { }`

2. **Same type params:** `func add(a, b int)` 

3. **Variadic:** `func f(nums ...int)` - must be last, pass slice with `...`

4. **Multiple returns:** Error always last, check with `if err != nil`

5. **Ignore returns:** Use `_` for values you don't need

6. **Functions are values:** Can assign to variables, pass as params, return

7. **Closures:** Inner functions can access and modify outer variables

8. **Shadowing:** `:=` creates new variable, `=` modifies existing

9. **Call by value:** Go copies everything (but maps/slices contain pointers)

10. **Slice + append:** Always `s = append(s, x)` and return if needed
