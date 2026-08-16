# Chapter 13: The Standard Library

---

## 1. io (Input/Output)

**Read from anywhere, Write to anywhere**

```go
// Read from file
file, _ := os.Open("data.txt")
data, _ := io.ReadAll(file)

// Read from string (same way!)
reader := strings.NewReader("hello")
data, _ := io.ReadAll(reader)
```

| Function | What it does |
|----------|--------------|
| `io.ReadAll(r)` | Read everything |
| `io.Copy(dst, src)` | Copy data |
| `io.WriteString(w, s)` | Write string |

---

## 2. time

```go
time.Now()                    // Current time
time.Sleep(2 * time.Second)   // Wait 2 seconds
time.Since(start)             // Duration since start
```

**Format (use reference: 2006-01-02 15:04:05):**
```go
now.Format("2006-01-02")      // 2024-08-16
now.Format("15:04:05")        // 14:30:45
```

**Timeout:**
```go
select {
case data := <-ch:
    fmt.Println(data)
case <-time.After(5 * time.Second):
    fmt.Println("Timeout!")
}
```

---

## 3. JSON

**Go → JSON (Marshal):**
```go
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

p := Person{Name: "John", Age: 25}
jsonData, _ := json.Marshal(p)
// {"name":"John","age":25}
```

**JSON → Go (Unmarshal):**
```go
jsonStr := `{"name":"Jane","age":30}`
var p Person
json.Unmarshal([]byte(jsonStr), &p)
```

**Struct Tags:**
```go
`json:"name"`       // Rename field
`json:"-"`          // Hide field
`json:",omitempty"` // Skip if empty
```

---

## 4. HTTP Client

```go
// GET request
resp, _ := http.Get("https://api.example.com")
body, _ := io.ReadAll(resp.Body)
resp.Body.Close()

// With timeout
client := &http.Client{Timeout: 10 * time.Second}
resp, _ := client.Get("https://api.example.com")
```

---

## 5. HTTP Server

```go
http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello!"))
})
http.ListenAndServe(":8080", nil)
```

**Request info:**
```go
r.Method       // GET, POST
r.URL.Path     // /hello
r.URL.Query()  // ?name=john
r.Header       // Headers
r.Body         // Body (io.Reader)
```

**Response:**
```go
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(200)
w.Write([]byte("data"))
```

---

## 6. ServeMux (Router)

```go
mux := http.NewServeMux()
mux.HandleFunc("/", homeHandler)
mux.HandleFunc("/api", apiHandler)
http.ListenAndServe(":8080", mux)
```

---

## 7. Middleware

**Code that runs before handler:**

```go
func Logger(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("Request:", r.URL.Path)
        next.ServeHTTP(w, r)  // Call actual handler
    })
}

// Use it
wrappedMux := Logger(mux)
http.ListenAndServe(":8080", wrappedMux)
```

---

## 8. slog (Structured Logging)

```go
slog.Info("message", "key1", value1, "key2", value2)
slog.Warn("warning")
slog.Error("error")
```

**JSON Logger:**
```go
handler := slog.NewJSONHandler(os.Stdout, nil)
logger := slog.New(handler)
logger.Info("user login", "name", "john")
// {"level":"INFO","msg":"user login","name":"john"}
```

---

## Quick Reference

| Package | Common Use |
|---------|------------|
| `io` | `io.ReadAll(r)` |
| `time` | `time.Now()`, `time.Sleep()` |
| `encoding/json` | `json.Marshal()`, `json.Unmarshal()` |
| `net/http` | `http.Get()`, `http.HandleFunc()` |
| `log/slog` | `slog.Info()`, `slog.Error()` |

---

## Server Flow

```
Request → Middleware → Handler → Response
            ↓
         (log it)
```
