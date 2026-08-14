# Chapter 10: Modules, Packages, and Imports

## Key Terms

| Term | What It Is |
|------|------------|
| Repository | Where code is stored (GitHub) |
| Module | One project (has go.mod) |
| Package | Folder with .go files |

---

## go.mod File

```bash
go mod init <module-name>
```

```go
module myproject

go 1.21

require (
    github.com/some/pkg v1.2.3
)
```

---

## Export = Uppercase

```go
func Double(x int) int { }  // ✅ Exported (Uppercase)
func helper(x int) int { }  // ❌ Not exported (lowercase)
```

---

## Project Structure

```
myproject/
├── go.mod
├── main.go          (package main)
├── math/
│   └── math.go      (package math)
└── utils/
    └── utils.go     (package utils)
```

---

## Import & Use

```go
import (
    "fmt"              // Standard library
    "myproject/math"   // Local package
)

func main() {
    math.Double(5)     // package.Function()
}
```

---

## Directives

| Directive | Use |
|-----------|-----|
| `replace` | Use fork/local copy |
| `exclude` | Block a version |
| `retract` | Mark your version as bad |

---

## Workspaces

```bash
go work init ./app
go work use ./library
```

Creates `go.work` (don't commit!)

---

## Key Rules

1. `go mod init` creates module
2. Uppercase = Exported (public)
3. lowercase = Not exported (private)
4. Import path = module + package
5. Must use what you import
