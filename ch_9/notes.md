# Chapter 9: Errors

## Creating Errors

```go
// Simple error
err := errors.New("something went wrong")

// Error with formatting
err := fmt.Errorf("cannot open %s", filename)
```

---

## Returning Errors

```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

// Always check!
result, err := divide(10, 0)
if err != nil {
    fmt.Println("Error:", err)
}
```

---

## Sentinel Errors

```go
var ErrNotFound = errors.New("not found")
var ErrInvalid = errors.New("invalid")

if err == ErrNotFound {
    // handle not found
}
```

---

## Custom Error Type

```go
type MyError struct {
    Code int
    Msg  string
}

func (e MyError) Error() string {
    return fmt.Sprintf("code %d: %s", e.Code, e.Msg)
}
```

---

## Wrapping Errors

```go
// %w keeps original error inside
return fmt.Errorf("in myFunc: %w", err)
```

---

## errors.Is (Check Value)

```go
if errors.Is(err, os.ErrNotExist) {
    fmt.Println("File not found!")
}
```

---

## errors.As (Check Type)

```go
var myErr *MyError
if errors.As(err, &myErr) {
    fmt.Println("Code:", myErr.Code)
}
```

---

## panic & recover

```go
// panic = crash
panic("fatal error!")

// recover = catch panic (must be in defer)
defer func() {
    if r := recover(); r != nil {
        fmt.Println("Recovered:", r)
    }
}()
```

---

## Quick Summary

| Concept | Use |
|---------|-----|
| `errors.New()` | Simple error |
| `fmt.Errorf("%w")` | Wrap error |
| `errors.Is()` | Check value |
| `errors.As()` | Check type |
| `panic` | Fatal only |
| `recover` | Catch panic |
