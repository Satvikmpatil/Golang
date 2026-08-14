package main

import (
	"fmt"
	"runtime"
)

// ============================================
// CHAPTER 11: Go Tooling
// ============================================

// --- 1. iota Enum (use with stringer) ---
type Direction int

const (
	North Direction = iota // 0
	South                  // 1
	East                   // 2
	West                   // 3
)

// To auto-generate String() method:
// 1. go install golang.org/x/tools/cmd/stringer@latest
// 2. Add: //go:generate stringer -type=Direction
// 3. Run: go generate ./...

// --- 2. Build Tags Example ---
// //go:build linux       → Only compile on Linux
// //go:build windows     → Only compile on Windows
// //go:build !windows    → Compile on everything except Windows
// //go:build ignore      → Never compile this file

func main() {
	// --- Runtime Info ---
	fmt.Println("=== Go Tooling Demo ===")
	fmt.Println()

	fmt.Println("--- 1. Runtime Info ---")
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Arch:", runtime.GOARCH)
	fmt.Println("Go Version:", runtime.Version())
	fmt.Println("Num CPU:", runtime.NumCPU())

	// --- iota/Enum Demo ---
	fmt.Println("\n--- 2. Enum (iota) ---")
	fmt.Println("North:", North)
	fmt.Println("South:", South)
	fmt.Println("East:", East)
	fmt.Println("West:", West)

	// --- Summary ---
	fmt.Println("\n--- Key Commands ---")
	fmt.Println("go build                    → Build")
	fmt.Println("go generate ./...           → Generate code")
	fmt.Println("GOOS=linux go build         → Cross-compile")
	fmt.Println("go version -m ./app         → See build info")
	fmt.Println("govulncheck ./...           → Security scan")
	fmt.Println("staticcheck ./...           → Find bugs")
}
