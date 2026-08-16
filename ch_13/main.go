package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"
)

// ============================================
// CHAPTER 13: The Standard Library
// ============================================

func main() {
	// =====================================
	// 1. io - READ & WRITE
	// =====================================
	fmt.Println("=== 1. io (Read/Write) ===")

	// Write to file
	file, _ := os.Create("test.txt")
	file.Write([]byte("Hello from Go!"))
	file.Close()

	// Read from file
	file2, _ := os.Open("test.txt")
	data, _ := io.ReadAll(file2)
	file2.Close()
	fmt.Println("File content:", string(data))

	// Read from string (same interface!)
	reader := strings.NewReader("Hello from string!")
	data2, _ := io.ReadAll(reader)
	fmt.Println("String content:", string(data2))

	// =====================================
	// 2. time - CLOCK & DURATION
	// =====================================
	fmt.Println("\n=== 2. time ===")

	// Current time
	now := time.Now()
	fmt.Println("Now:", now)

	// Format time
	fmt.Println("Date:", now.Format("2006-01-02"))
	fmt.Println("Time:", now.Format("15:04:05"))

	// Measure duration
	start := time.Now()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Took:", time.Since(start))

	// =====================================
	// 3. JSON - MARSHAL (Go → JSON)
	// =====================================
	fmt.Println("\n=== 3. JSON Marshal ===")

	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	p := Person{Name: "John", Age: 25}
	jsonBytes, _ := json.Marshal(p)
	fmt.Println("JSON:", string(jsonBytes))

	// =====================================
	// 4. JSON - UNMARSHAL (JSON → Go)
	// =====================================
	fmt.Println("\n=== 4. JSON Unmarshal ===")

	jsonStr := `{"name":"Jane","age":30}`
	var p2 Person
	json.Unmarshal([]byte(jsonStr), &p2)
	fmt.Println("Name:", p2.Name)
	fmt.Println("Age:", p2.Age)

	// =====================================
	// 5. JSON - HIDE FIELDS
	// =====================================
	fmt.Println("\n=== 5. JSON Hide Fields ===")

	type User struct {
		Name     string `json:"name"`
		Password string `json:"-"` // Hidden!
	}

	u := User{Name: "John", Password: "secret123"}
	userJSON, _ := json.Marshal(u)
	fmt.Println("User JSON:", string(userJSON)) // No password!

	// =====================================
	// 6. HTTP CLIENT - GET REQUEST
	// =====================================
	fmt.Println("\n=== 6. HTTP Client ===")

	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Println("Response length:", len(body), "bytes")
	}

	// =====================================
	// 7. SLOG - STRUCTURED LOGGING
	// =====================================
	fmt.Println("\n=== 7. slog (Structured Logging) ===")

	slog.Info("user logged in", "name", "john", "id", 123)
	slog.Warn("disk space low", "percent", 85)
	slog.Error("database error", "code", 500)

	// JSON logger
	fmt.Println("\n--- JSON Logger ---")
	jsonHandler := slog.NewJSONHandler(os.Stdout, nil)
	jsonLogger := slog.New(jsonHandler)
	jsonLogger.Info("json log", "user", "jane", "action", "purchase")

	// =====================================
	// 8. HTTP SERVER
	// =====================================
	fmt.Println("\n=== 8. HTTP Server ===")
	fmt.Println("Starting server on :8080...")
	fmt.Println("Try: http://localhost:8080/")
	fmt.Println("Try: http://localhost:8080/hello")
	fmt.Println("Try: http://localhost:8080/users")
	fmt.Println("Try: http://localhost:8080/time")
	fmt.Println("\nPress Ctrl+C to stop")

	// Create router
	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/users", usersHandler)
	mux.HandleFunc("/time", timeHandler)

	// Wrap with logger middleware
	loggedMux := LoggerMiddleware(mux)

	// Start server
	http.ListenAndServe(":8080", loggedMux)
}

// =====================================
// HANDLERS
// =====================================

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome Home!"))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	users := []map[string]any{
		{"id": 1, "name": "John"},
		{"id": 2, "name": "Jane"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now().Format(time.RFC3339)
	w.Write([]byte("Current time: " + now))
}

// =====================================
// MIDDLEWARE
// =====================================

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		slog.Info("request started", "method", r.Method, "path", r.URL.Path)

		next.ServeHTTP(w, r)

		slog.Info("request completed", "duration", time.Since(start))
	})
}
