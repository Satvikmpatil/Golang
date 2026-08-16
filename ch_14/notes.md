# Chapter 14: The Context

---

## What is Context?

**Context = Remote control for goroutines**

- Stop running tasks
- Set timeouts
- Pass data

---

## 1. Create Context

```go
ctx := context.Background()  // Starting point (empty)
```

---

## 2. Cancel Manually

```go
ctx, cancel := context.WithCancel(context.Background())

go func() {
    for {
        select {
        case <-ctx.Done():  // Cancelled?
            return          // Stop!
        default:
            // Keep working...
        }
    }
}()

cancel()  // Stop the goroutine!
```

---

## 3. Timeout (Auto Cancel)

```go
// Cancel after 5 seconds
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()  // Always call!

select {
case result := <-doWork():
    fmt.Println(result)
case <-ctx.Done():
    fmt.Println("Timeout!")
}
```

---

## 4. Deadline (Cancel at specific time)

```go
deadline := time.Now().Add(10 * time.Second)
ctx, cancel := context.WithDeadline(context.Background(), deadline)
defer cancel()
```

---

## 5. Pass Values

```go
// Store
ctx := context.WithValue(context.Background(), "userID", 123)

// Get
userID := ctx.Value("userID").(int)
```

---

## 6. HTTP Request with Timeout

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

req, _ := http.NewRequestWithContext(ctx, "GET", "https://api.com", nil)
resp, err := http.DefaultClient.Do(req)
```

---

## 7. Check Why Cancelled

```go
err := ctx.Err()

if err == context.Canceled {
    // Someone called cancel()
}

if err == context.DeadlineExceeded {
    // Timeout happened
}
```

---

## Quick Reference

| Function | What it does |
|----------|--------------|
| `context.Background()` | Create empty context |
| `context.WithCancel(ctx)` | Add manual cancel |
| `context.WithTimeout(ctx, dur)` | Add timeout |
| `context.WithDeadline(ctx, time)` | Add deadline |
| `context.WithValue(ctx, k, v)` | Add value |

| Method | What it does |
|--------|--------------|
| `ctx.Done()` | Channel that closes on cancel |
| `ctx.Err()` | Error (why cancelled) |
| `ctx.Value(key)` | Get stored value |
| `ctx.Deadline()` | Get timeout time |

---

## The Pattern

```go
func doWork(ctx context.Context) error {
    for {
        select {
        case <-ctx.Done():
            return ctx.Err()  // Stop!
        default:
            // Work...
        }
    }
}
```

---

## Rules

1. Context is **first parameter**: `func foo(ctx context.Context, ...)`
2. Always call `cancel()` (use `defer cancel()`)
3. Pass context to HTTP requests, DB calls
4. Check `ctx.Done()` in long loops

---

## Simple Analogy

| Context | Real Life |
|---------|-----------|
| `WithCancel` | Walkie-talkie "STOP" |
| `WithTimeout` | Microwave timer |
| `WithValue` | Passing a note |
| `ctx.Done()` | Listening for "STOP" |
