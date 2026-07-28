# Chapter 6: Pointers (Simple Guide)

## What is a Pointer?

**Pointer = Address of a variable (where it lives in memory)**

Think of it like:
```
Variable = House
Pointer  = Address of that house
```

---

## Two Symbols to Remember

| Symbol | Name | What It Does |
|--------|------|--------------|
| `&` | Address of | "Where is this?" |
| `*` | Value at | "What's there?" |

```go
x := 10
ptr := &x         // ptr = address of x

fmt.Println(x)    // 10 (the value)
fmt.Println(ptr)  // 0xc0000... (the address)
fmt.Println(*ptr) // 10 (value AT that address)
```

---

## Change Value Through Pointer

```go
x := 10
ptr := &x

*ptr = 50         // Put 50 at that address

fmt.Println(x)    // 50 (x changed!)
```

---

## Pointer Type

```go
var x int = 10
var ptr *int      // *int = "pointer to int"
ptr = &x

// Common types:
// *int    = pointer to int
// *string = pointer to string
// *bool   = pointer to bool
// *Foo    = pointer to Foo struct
```

---

## new() Function

```go
ptr := new(int)   // Creates pointer to zero-value

fmt.Println(ptr)  // 0xc0000... (some address)
fmt.Println(*ptr) // 0 (zero value of int)
```

---

## nil Pointer (DANGER!)

```go
var ptr *int      // ptr = nil (points to nothing)

fmt.Println(ptr == nil)  // true

// This will CRASH your program!
// fmt.Println(*ptr)      // PANIC!

// Always check first:
if ptr != nil {
    fmt.Println(*ptr)     // Safe!
}
```

---

## Can't Get Address of Literals

```go
// These DON'T work:
ptr := &10        // Error!
ptr := &"hello"   // Error!

// Solution: Use variable first
x := 10
ptr := &x         // Works!
```

---

## Helper Function for Literals

```go
func makePointer[T any](t T) *T {
    return &t
}

// Now you can do:
ptr := makePointer(10)       // *int pointing to 10
ptr := makePointer("hello")  // *string pointing to "hello"
```

---

## Struct with Pointer Field

```go
type Person struct {
    Name   string
    Age    *int    // Pointer = can be nil
}

// Problem: Can't do this
p := Person{
    Name: "Bob",
    Age:  &25,     // Error! Can't &literal
}

// Solution 1: Variable
age := 25
p := Person{Name: "Bob", Age: &age}

// Solution 2: Helper function
p := Person{Name: "Bob", Age: makePointer(25)}
```

---

## Functions: Value vs Pointer

### Without Pointer (Can't change original)

```go
func change(n int) {
    n = 100       // Changes copy only
}

x := 10
change(x)
fmt.Println(x)    // 10 (NOT changed)
```

### With Pointer (Can change original)

```go
func change(ptr *int) {
    *ptr = 100    // Changes original
}

x := 10
change(&x)
fmt.Println(x)    // 100 (CHANGED!)
```

---

## 3 Rules of Pointers in Functions

### Rule 1: Change field via pointer = Works!

```go
type Foo struct { x int }

func changeField(f *Foo) {
    f.x = 20      // Changes original!
}

obj := &Foo{x: 10}
changeField(obj)
fmt.Println(obj.x)  // 20
```

### Rule 2: Reassign pointer = Doesn't work!

```go
func reassign(f *Foo) {
    f = &Foo{x: 30}   // Only changes copy!
}

obj := &Foo{x: 10}
reassign(obj)
fmt.Println(obj.x)  // Still 10!
```

### Rule 3: nil pointer = Can't modify!

```go
func tryModify(f *Foo) {
    f = &Foo{x: 40}   // Can't change nil to non-nil
}

var obj *Foo = nil
tryModify(obj)
fmt.Println(obj)    // Still nil!
```

### Visual Explanation

```
Rule 1: f.x = 20
┌─────────┐
│ obj ────┼──→ [x: 10] ──→ [x: 20]  Both see change!
│ f ──────┼──↗
└─────────┘

Rule 2: f = &Foo{x: 30}
┌─────────┐
│ obj ────┼──→ [x: 10]   obj still points here
│ f ──────┼──→ [x: 30]   f points to NEW box
└─────────┘

Rule 3: nil → can't change
┌─────────┐
│ obj = nil
│ f = nil  ──→ f = &Foo{}  (only f changes)
└─────────┘
```

---

## Right Way vs Wrong Way

```go
// WRONG: Changing pointer itself
func failedUpdate(px *int) {
    x2 := 20
    px = &x2      // Changes copy of pointer!
}

// RIGHT: Changing value at pointer
func update(px *int) {
    *px = 20      // Changes actual value!
}

x := 10
failedUpdate(&x)
fmt.Println(x)    // 10 (unchanged)

update(&x)
fmt.Println(x)    // 20 (changed!)
```

**Remember:**
```
px = &x2   → Changes WHERE pointer points (copy only)
*px = 20   → Changes WHAT'S at address (works!)
```

---

## Pointers Are Last Resort

### Don't Do This:

```go
func MakeFoo(f *Foo) error {
    f.Field1 = "val"
    f.Field2 = 20
    return nil
}
```

### Do This Instead:

```go
func MakeFoo() (Foo, error) {
    return Foo{
        Field1: "val",
        Field2: 20,
    }, nil
}
```

**Rule:** Return values, don't fill pointers.

---

## When to Use Pointers

| Use Pointer | Use Value |
|-------------|-----------|
| Need to modify original | Most cases! |
| Very large struct (10MB+) | Small/medium data |
| Interface requires it (JSON) | Simple functions |
| Indicate "no value" (nil) | Known values |

---

## Pointer Performance

| Data Size | Faster Option |
|-----------|---------------|
| Small (< 10KB) | Value |
| Medium (10KB - 10MB) | Either |
| Large (> 10MB) | Pointer |

**For most code:** Use values. Don't worry about performance.

---

## Zero Value vs No Value

```go
type Person struct {
    Name string
    Age  int      // Zero = 0, can't tell if "not set"
}

type PersonBetter struct {
    Name string
    Age  *int     // nil = "not set", &0 = "is zero"
}

// Check if age was provided:
if p.Age == nil {
    fmt.Println("Age not provided")
} else {
    fmt.Println("Age is", *p.Age)
}
```

---

## Maps vs Slices in Functions

### Map = Always changes original

```go
func modifyMap(m map[string]int) {
    m["a"] = 100
}

myMap := map[string]int{"a": 1}
modifyMap(myMap)
fmt.Println(myMap["a"])  // 100 (changed!)
```

**Why?** Map is internally a pointer.

### Slice = Tricky!

```go
func modifySlice(s []int) {
    s[0] = 100        // Changes original!
    s = append(s, 5)  // Doesn't change original!
}

mySlice := []int{1, 2, 3}
modifySlice(mySlice)
fmt.Println(mySlice)  // [100 2 3] (only s[0] changed)
```

### Why Slice is Tricky

```
Slice = struct with 3 fields:
┌────────┬────────┬────────┐
│ length │capacity│ pointer│──→ [actual data]
└────────┴────────┴────────┘

Function gets COPY of this struct.
- s[0] = 100  → Uses pointer → Changes data!
- append()    → May change pointer → Only copy updated!
```

### Visual: append Failure

```
Before append:
Original: [len=3, cap=3, ptr]──→ [1, 2, 3]
Copy:     [len=3, cap=3, ptr]──↗

After append (no room, new memory):
Original: [len=3, cap=3, ptr]──→ [1, 2, 3]     (unchanged!)
Copy:     [len=4, cap=6, ptr]──→ [1, 2, 3, 5]  (new array)
```

---

## Slices as Buffers

### What is a Buffer?

**Buffer = A reusable container for data**

Think of it like:
```
Buffer = A reusable plate
- You eat food from plate
- Wash plate
- Use same plate again

vs

New plate every time = wasteful!
```

### Real World Example

Imagine you're reading a book and taking notes:

**Bad Way (wasteful):**
```
Read page 1 → Buy new notebook → Write notes → Throw notebook
Read page 2 → Buy new notebook → Write notes → Throw notebook
Read page 3 → Buy new notebook → Write notes → Throw notebook
...
```
You bought 100 notebooks for 100 pages! Wasteful!

**Good Way (buffer):**
```
Buy ONE notebook
Read page 1 → Write in notebook → Erase
Read page 2 → Write in notebook → Erase
Read page 3 → Write in notebook → Erase
...
```
You used ONE notebook for all pages! Efficient!

### Code Example

**Bad: Creates new memory every loop**

```go
for i := 0; i < 1000; i++ {
    data := make([]byte, 100)    // New memory EVERY time!
    file.Read(data)
    process(data)
}
// Created 1000 slices! All become garbage!
```

**Good: Reuse same memory**

```go
buffer := make([]byte, 100)  // Create ONCE!

for i := 0; i < 1000; i++ {
    count, err := file.Read(buffer)  // Reuse same memory
    process(buffer[:count])
    if err == io.EOF {
        break
    }
}
// Created only 1 slice! Much better!
```

### Why Buffer is Better

```
Bad way:
Loop 1: [new memory] → garbage
Loop 2: [new memory] → garbage
Loop 3: [new memory] → garbage
...
Result: 1000 pieces of garbage for GC to clean!

Good way (buffer):
Loop 1: [same memory] ← reuse
Loop 2: [same memory] ← reuse
Loop 3: [same memory] ← reuse
...
Result: 0 garbage! Fast!
```

---

## Stack vs Heap

### Think of it Like a House

```
Stack = Your desk (fast, organized, limited space)
Heap  = Your garage (slower, messy, lots of space)
```

### Stack - The Fast Storage

**What is Stack?**
- Fast memory
- Each function gets its own space
- When function ends, memory is automatically freed

**Example:**
```go
func add(a, b int) int {
    result := a + b    // result is on STACK
    return result
}                      // Function ends, stack is cleared!
```

**Visual:**
```
Function starts:
┌─────────────────┐
│ Stack           │
│ a = 5           │
│ b = 3           │
│ result = 8      │
└─────────────────┘

Function ends:
┌─────────────────┐
│ Stack           │
│ (empty - freed!)│
└─────────────────┘
```

### Heap - The Slow Storage

**What is Heap?**
- Slower memory
- Data can live beyond function
- Garbage Collector must clean it

**When data goes to Heap:**
```go
func createUser() *User {
    u := User{Name: "Bob"}  // Must go to HEAP!
    return &u               // Because pointer is returned
}
// If u was on stack, it would be destroyed when function ends!
// So Go puts it on heap to keep it alive.
```

### Stack vs Heap Comparison

| Feature | Stack | Heap |
|---------|-------|------|
| Speed | Very Fast | Slower |
| Size limit | Small (1-8 MB) | Large (GBs) |
| Cleanup | Automatic (instant) | Garbage Collector |
| When used | Local variables | Pointers that escape |

### What Goes Where?

**Stack (fast):**
```go
x := 10           // int → stack
name := "hello"   // string → stack
arr := [3]int{}   // array → stack
p := Person{}     // struct → stack
```

**Heap (slower):**
```go
ptr := &x         // data ptr points to might go to heap
slice := make([]int, 1000)  // large data → heap
ptr := new(Person)          // new() → heap
```

### Simple Rule

```
Local variable, stays in function → Stack
Pointer returned from function    → Heap
```

---

## Garbage Collector

### What is Garbage?

**Garbage = Memory that no one is using anymore**

Think of it like:
```
You have a toy.
You throw toy in corner.
You forget about toy.
Toy = garbage!

Mom comes and cleans it up.
Mom = Garbage Collector!
```

### Code Example

```go
func createGarbage() {
    x := &Foo{Name: "test"}   // x points to Foo
    fmt.Println(x.Name)
}                              // Function ends
                               // x is gone
                               // Foo has no pointer → GARBAGE!
                               // GC will clean it
```

### Visual

```
BEFORE function ends:
┌─────────┐
│ x ──────┼──→ [Foo{Name: "test"}]
└─────────┘
             ↑ Has pointer, NOT garbage

AFTER function ends:
┌─────────┐
│ (x gone)│    [Foo{Name: "test"}]  ← No pointer!
└─────────┘
                ↑ GARBAGE! GC will clean
```

### When is GC Called?

```
You don't call GC manually!
Go automatically runs GC when:
- Heap gets too full
- Based on GOGC setting
```

### GC Takes Time

```
Your program running: ████████████████████
                           ↑
                      GC runs here
                      (your program pauses briefly)

More garbage = More GC runs = Slower program!
```

### How to Create Less Garbage

| Do This | Why |
|---------|-----|
| Use values, not pointers | Stays on stack, no GC needed |
| Reuse slices (buffers) | No new allocations |
| `[]Person` not `[]*Person` | Data together, faster |
| Preallocate slices | `make([]int, 0, 100)` |

### Example: Good vs Bad

```go
// BAD: Creates garbage
func bad() {
    for i := 0; i < 1000; i++ {
        p := &Person{Name: "test"}  // 1000 allocations!
        fmt.Println(p.Name)
    }
}
// 1000 Person objects become garbage!

// GOOD: No garbage
func good() {
    p := Person{}  // 1 allocation
    for i := 0; i < 1000; i++ {
        p.Name = "test"  // Reuse same struct
        fmt.Println(p.Name)
    }
}
// 0 garbage!
```

---

## GOGC and GOMEMLIMIT

### What is GOGC?

**GOGC = "How often should GC run?"**

Think of it like:
```
Your room gets messy.
When do you clean?

GOGC=100 → Clean when mess DOUBLES
GOGC=50  → Clean when mess is HALF more
GOGC=200 → Clean when mess TRIPLES
```

### GOGC Values

| GOGC | Meaning | Effect |
|------|---------|--------|
| `100` (default) | GC when heap doubles | Balanced |
| `50` | GC more often | Less memory, slower |
| `200` | GC less often | More memory, faster |
| `off` | Never GC | Dangerous! Memory grows forever |

### How to Set GOGC

```bash
# In terminal before running program:
GOGC=200 go run main.go

# Or in code:
import "runtime/debug"
debug.SetGCPercent(200)
```

### What is GOMEMLIMIT?

**GOMEMLIMIT = "Maximum memory my program can use"**

Think of it like:
```
Your room is 10 square meters.
You can't have more stuff than fits in room!
```

### GOMEMLIMIT Values

```bash
GOMEMLIMIT=512MiB ./myprogram   # Max 512 MB
GOMEMLIMIT=1GiB ./myprogram     # Max 1 GB
GOMEMLIMIT=2GiB ./myprogram     # Max 2 GB
```

### When to Use These?

| Setting | When to Use |
|---------|-------------|
| Default GOGC | Most programs, don't change |
| Higher GOGC | Need speed, have extra memory |
| Lower GOGC | Limited memory (small server) |
| GOMEMLIMIT | Docker/container with memory limit |

### Simple Advice

```
For beginners: Don't touch these settings!
Go's defaults work great for 99% of programs.
Only change if you have specific memory problems.
```

---

## Quick Summary Table

| Concept | Remember |
|---------|----------|
| `&x` | Get address of x |
| `*ptr` | Get value at ptr |
| `*int` | Type: pointer to int |
| `new(int)` | Create pointer to zero-value |
| `nil` pointer | Points to nothing, panic if used |
| `*ptr = 20` | Changes original (works!) |
| `ptr = &x` | Changes copy only (doesn't work!) |
| Map in function | Always changes original |
| Slice `s[0]=x` | Changes original |
| Slice `append` | Doesn't change original |
| Buffer | Reuse slice to avoid garbage |
| Stack | Fast, auto cleanup |
| Heap | Slow, GC cleans |
| Garbage | Data with no pointers |
| GOGC | How often GC runs |
| GOMEMLIMIT | Max memory allowed |

---

## Golden Rules

1. **Use values by default** - simpler, safer, faster

2. **Use pointers only when needed:**
   - Modify original data
   - Very large structs (10MB+)
   - Need nil to mean "no value"

3. **Always check nil before `*ptr`**

4. **`*ptr = x` changes data, `ptr = &x` doesn't**

5. **Return values, don't fill pointer params**

6. **Slices: can change content, can't change size**

7. **Reuse buffers to reduce garbage**

8. **Less pointers = less heap = less GC = faster!**

---

## Common Mistakes

```go
// Mistake 1: Dereferencing nil
var p *int
fmt.Println(*p)  // PANIC! Always check nil first!

// Fix:
if p != nil {
    fmt.Println(*p)
}
```

```go
// Mistake 2: Expecting append to change original
func addItem(s []int) {
    s = append(s, 5)  // Original unchanged!
}

// Fix: Return the slice
func addItem(s []int) []int {
    return append(s, 5)
}
s = addItem(s)  // Assign back!
```

```go
// Mistake 3: Reassigning pointer param
func update(p *int) {
    x := 20
    p = &x  // Doesn't work!
}

// Fix: Change value at pointer
func update(p *int) {
    *p = 20  // Works!
}
```

```go
// Mistake 4: Using pointer when value is fine
func getName(p *Person) string {  // Bad - why pointer?
    return p.Name
}

// Fix: Use value (unless you need to modify)
func getName(p Person) string {   // Good!
    return p.Name
}
```

```go
// Mistake 5: Creating garbage in loop
for i := 0; i < 1000; i++ {
    data := make([]byte, 100)  // 1000 allocations!
}

// Fix: Reuse buffer
data := make([]byte, 100)  // 1 allocation
for i := 0; i < 1000; i++ {
    // use data
}
```

---

## Key Takeaways

1. **Pointer** = Address where data lives

2. **`&`** = "where is it?" / **`*`** = "what's there?"

3. **Pointers** let you modify original data

4. **nil pointer** = nothing, will crash if used

5. **Prefer values** over pointers (simpler, faster!)

6. **Map** always shared, **Slice** tricky with append

7. **Stack** = fast (local vars), **Heap** = slow (pointers)

8. **Garbage** = unused memory, GC cleans it

9. **Less garbage** = faster program

10. **Buffer** = reuse slice to avoid allocations
