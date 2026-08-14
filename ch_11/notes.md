# Chapter 11: Go Tooling

## Key Commands

| Command | Purpose |
|---------|---------|
| `go build` | Compile |
| `go generate ./...` | Auto-generate code |
| `go version -m app` | See build info |
| `govulncheck ./...` | Security scan |
| `staticcheck ./...` | Find bugs |

---

## Cross-Compilation

```bash
# Linux 64-bit
GOOS=linux GOARCH=amd64 go build

# Windows
GOOS=windows GOARCH=amd64 go build

# Mac ARM (M1/M2)
GOOS=darwin GOARCH=arm64 go build
```

| GOOS | OS |
|------|----|
| `linux` | Linux |
| `windows` | Windows |
| `darwin` | macOS |

| GOARCH | CPU |
|--------|-----|
| `amd64` | 64-bit Intel |
| `arm64` | ARM 64-bit |
| `386` | 32-bit |

---

## Build Tags

```go
//go:build linux
package main
// Only compiles on Linux
```

```go
//go:build windows && amd64
// Only Windows 64-bit
```

```go
//go:build ignore
// Never compiles
```

---

## go generate

```go
//go:generate stringer -type=Direction

type Direction int
const (
    North Direction = iota
    South
)
```

```bash
go install golang.org/x/tools/cmd/stringer@latest
go generate ./...
```

---

## Useful Tools

```bash
# Install
go install golang.org/x/vuln/cmd/govulncheck@latest
go install honnef.co/go/tools/cmd/staticcheck@latest

# Use
govulncheck ./...
staticcheck ./...
```

---

## Build Info

```bash
go version -m ./myapp
```

Shows: Go version, dependencies, build settings

---

## Summary

```
go build              → Compile
go generate           → Generate code
GOOS=linux go build   → Cross-compile
go version -m app     → Build info
//go:build linux      → Build tag
```
