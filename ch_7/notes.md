# Chapter 7: Types, Methods, and Interfaces

## 1. User-Defined Types

```go
type Score int                    // Score is just an int
type Person struct { ... }        // Struct type
type Converter func(string) int   // Function type
type TeamScores map[string]int    // Map type
```

---

## 2. Methods

**Method = Function attached to a type**

```go
type Person struct {
    Name string
    Age  int
}

func (p Person) Greet() string {
    return "Hello, " + p.Name
}

person := Person{Name: "Bob"}
person.Greet()  // "Hello, Bob"
```

---

## 3. Value vs Pointer Receiver

| | Value `(p Person)` | Pointer `(p *Person)` |
|---|---|---|
| Modifies? | ❌ No | ✅ Yes |
| Use when | Just reading | Changing data |

```go
// Value - doesn't modify
func (c Counter) Total() int {
    return c.total
}

// Pointer - modifies!
func (c *Counter) Increment() {
    c.total++
}
```

---

## 4. nil Receiver

```go
func (t *Tree) Insert(val int) *Tree {
    if t == nil {
        return &Tree{val: val}  // Handle nil!
    }
    // ...
}

var t *Tree  // nil
t = t.Insert(5)  // Works!
```

---

## 5. iota (Enums)

```go
type Day int

const (
    _         Day = iota  // 0 (skip)
    Monday                // 1
    Tuesday               // 2
    Wednesday             // 3
)
```

---

## 6. Embedding (Composition)

```go
type Employee struct {
    Name string
    ID   string
}

type Manager struct {
    Employee        // Embedded!
    Reports []Employee
}

m := Manager{Employee: Employee{Name: "Bob"}}
m.Name  // Direct access! (not m.Employee.Name)
```

**NOT inheritance! Just composition.**

---

## 7. Interfaces

**Interface = List of methods**

```go
type Speaker interface {
    Speak()
}

type Dog struct{}
func (d Dog) Speak() { fmt.Println("Woof!") }

type Cat struct{}
func (c Cat) Speak() { fmt.Println("Meow!") }

// One function for all!
func MakeSound(s Speaker) {
    s.Speak()
}

MakeSound(Dog{})  // Woof!
MakeSound(Cat{})  // Meow!
```

**No "implements" keyword! Just have the methods.**

---

## Quick Summary

| Concept | Remember |
|---------|----------|
| Method | Function on type |
| Value receiver | Doesn't modify |
| Pointer receiver | Can modify |
| iota | Auto: 0, 1, 2, 3... |
| Embedding | Composition |
| Interface | List of methods |

---

## Key Rules

1. Pointer receiver = can modify
2. Go has NO inheritance
3. Embedding = composition
4. Interface = "what can you do?"
5. No "implements" keyword needed
