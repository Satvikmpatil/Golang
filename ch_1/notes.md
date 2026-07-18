# Chapter 1: Setting Up Your Go Environment

## Why Go is Special

- Go compiles to **one single file** (binary)
- No need to install anything on other computers to run it
- You: need Go to BUILD
- Friend: does NOT need Go to RUN

---

## Creating a Go Project

```bash
mkdir ch1           # Create folder
cd ch1              # Go inside
go mod init hello   # Make it a Go project
```

This creates `go.mod` file (project ID card):
```
module hello
go 1.20
```

---

## Basic Go File Structure

```go
package main         // Entry point of program

import "fmt"         // Import packages you need

func main() {        // Program starts here
    fmt.Println("Hello, world!")
}
```

| Part | Meaning |
|------|---------|
| `package main` | This is where program starts |
| `import "fmt"` | I want to use fmt package |
| `func main()` | The starting function |

---

## Go Formatting Rules

1. Use **tabs** for indentation (not spaces)
2. Opening brace `{` must be on **same line**
3. Go adds semicolons automatically

**WRONG:**
```go
func main()
{              // ERROR!
}
```

**CORRECT:**
```go
func main() {  // Brace on same line
}
```

---

## Essential Commands

| Command | What it does |
|---------|--------------|
| `go build` | Compile your code |
| `go run file.go` | Compile + run immediately |
| `go fmt ./...` | Auto-format your code |
| `go vet ./...` | Find hidden bugs |

---

## go vet Example

```go
fmt.Printf("Hello, %s!\n")      // Bug: %s needs a value!
```

```bash
$ go vet ./...
# Error: %s needs arg #1, but got 0 args
```

**Fix:**
```go
fmt.Printf("Hello, %s!\n", "world")
```

---

## Makefile (Automate Everything)

Create a file called `Makefile`:

```makefile
.DEFAULT_GOAL := build

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build
```

Now just run:
```bash
make         # Runs fmt → vet → build
make fmt     # Only format
make vet     # Format + check bugs
```

---

## Tools for Writing Go

| Tool | Cost | Best for |
|------|------|----------|
| VS Code | Free | Most developers |
| GoLand | Paid | Professional use |
| Go Playground | Free | Quick testing online |

**Go Playground:** [play.golang.org](https://play.golang.org)
- Test code in browser
- No installation needed
- Don't put passwords there!

---

## Quick Reference

```bash
# Start new project
mkdir myproject && cd myproject
go mod init myproject

# Daily workflow
go fmt ./...     # Format code
go vet ./...     # Check for bugs
go build         # Compile
go run main.go   # Run directly

# Or just use Makefile
make
```

---

## Key Takeaways

1. Go compiles to standalone binary - easy to share
2. Every project needs `go.mod` file
3. Every program needs `package main` + `func main()`
4. Always run `go fmt` and `go vet` before committing
5. Brace `{` must be on same line - it's not optional!
