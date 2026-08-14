package utils

import "fmt"

// Greet returns a greeting (Exported)
func Greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// Max returns the larger number (Exported)
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// privateFunc is not exported (lowercase)
func privateFunc() string {
	return "can't use from outside"
}
