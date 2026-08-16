# Chapter 12: Concurrency in Go

## What is Concurrency?

**Concurrency = Running multiple tasks at the same time**

---

## 1. Goroutine

**Run function in background**

```go
// Normal (waits)
doWork()

// Goroutine (runs in background)
go doWork()
```

---

## 2. Channel

**Send data between goroutines**

```go
ch := make(chan string)   // Create

ch <- "hello"             // Send
msg := <-ch               // Receive
```

---

## 3. Buffered Channel

```go
ch := make(chan int, 3)   // Can hold 3 values

ch <- 1  // OK
ch <- 2  // OK
ch <- 3  // OK
// ch <- 4  // Blocks! Buffer full
```

---

## 4. Select

**Listen to multiple channels**

```go
select {
case msg := <-ch1:
    fmt.Println(msg)
case msg := <-ch2:
    fmt.Println(msg)
case <-time.After(time.Second):
    fmt.Println("Timeout!")
}
```

---

## 5. Mutex

**Protect shared data**

```go
var mu sync.Mutex
var counter int

mu.Lock()
counter++
mu.Unlock()
```

---

## 6. WaitGroup

**Wait for goroutines to finish**

```go
var wg sync.WaitGroup

wg.Add(1)          // Add task
go func() {
    defer wg.Done() // Mark done
    // work
}()
wg.Wait()          // Wait for all
```

---

## 7. Close Channel

```go
close(ch)                 // Close channel

for val := range ch {     // Loop until closed
    fmt.Println(val)
}
```

---

## Quick Reference

| What | Code |
|------|------|
| Goroutine | `go func()` |
| Create channel | `make(chan int)` |
| Send | `ch <- value` |
| Receive | `<-ch` |
| Buffered | `make(chan int, 5)` |
| Select | `select { case... }` |
| Mutex | `mu.Lock()` / `mu.Unlock()` |
| WaitGroup | `Add()` / `Done()` / `Wait()` |
| Close | `close(ch)` |

---

## When to Use What?

| Need | Use |
|------|-----|
| Run in background | Goroutine |
| Pass data | Channel |
| Multiple channels | Select |
| Protect variable | Mutex |
| Wait for finish | WaitGroup |

---

## Key Rules

1. `go func()` runs in background
2. Channels block until send/receive
3. Close channels from sender
4. Use `defer mu.Unlock()` after Lock
5. Prefer channels over mutex
