# Chapter 15: Writing Tests

---

## Basic Test

**File:** `xxx_test.go`
**Function:** `TestXxx(t *testing.T)`

```go
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Error("Expected 5, got", result)
    }
}
```

---

## Table Test

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        a, b     int
        expected int
    }{
        {1, 1, 2},
        {2, 3, 5},
    }
    
    for _, tc := range tests {
        if Add(tc.a, tc.b) != tc.expected {
            t.Error("wrong result")
        }
    }
}
```

---

## Subtests

```go
t.Run("test name", func(t *testing.T) {
    // test code
})
```

---

## Error Methods

| Method | Does |
|--------|------|
| `t.Error()` | Log error, continue |
| `t.Errorf()` | Log formatted, continue |
| `t.Fatal()` | Log error, STOP |
| `t.Fatalf()` | Log formatted, STOP |

---

## HTTP Test

```go
func TestHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/path", nil)
    rr := httptest.NewRecorder()
    
    MyHandler(rr, req)
    
    if rr.Code != 200 {
        t.Error("bad status")
    }
}
```

---

## Benchmark

```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(2, 3)
    }
}
```

---

## Commands

| Command | What it does |
|---------|--------------|
| `go test` | Run tests |
| `go test -v` | Verbose |
| `go test -cover` | Coverage % |
| `go test -race` | Find race bugs |
| `go test -bench=.` | Benchmarks |
| `go test -run TestAdd` | Run specific test |

---

## Rules

1. File ends with `_test.go`
2. Function starts with `Test`
3. Takes `*testing.T`
4. Benchmark starts with `Benchmark`
5. Benchmark takes `*testing.B`

---

## Coverage

```bash
go test -coverprofile=cover.out
go tool cover -html=cover.out
```

---

## Quick Example

```go
// math.go
func Add(a, b int) int {
    return a + b
}

// math_test.go
func TestAdd(t *testing.T) {
    if Add(2, 3) != 5 {
        t.Error("Add failed")
    }
}
```
