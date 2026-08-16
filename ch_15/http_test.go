package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// =====================================
// 8. HTTP HANDLER TEST
// =====================================

// Handler to test
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
}

func TestHelloHandler(t *testing.T) {
	// Create fake request
	req := httptest.NewRequest("GET", "/hello", nil)

	// Create fake response recorder
	rr := httptest.NewRecorder()

	// Call handler
	HelloHandler(rr, req)

	// Check status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rr.Code)
	}

	// Check body
	expected := "Hello!"
	if rr.Body.String() != expected {
		t.Errorf("Expected %q, got %q", expected, rr.Body.String())
	}
}

// Handler with query parameter
func GreetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	w.Write([]byte("Hello, " + name + "!"))
}

func TestGreetHandler(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		expected string
	}{
		{"with name", "/greet?name=John", "Hello, John!"},
		{"no name", "/greet", "Hello, World!"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tc.url, nil)
			rr := httptest.NewRecorder()

			GreetHandler(rr, req)

			if rr.Body.String() != tc.expected {
				t.Errorf("got %q, want %q", rr.Body.String(), tc.expected)
			}
		})
	}
}
