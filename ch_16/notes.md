# Chapter 16: Reflect, Unsafe, and Cgo

⚠️ **WARNING: Advanced topics. Avoid unless necessary!**

---

## 1. Reflect

**Inspect types at runtime**

```go
import "reflect"

x := 42
reflect.TypeOf(x)   // int
reflect.ValueOf(x)  // 42
reflect.TypeOf(x).Kind()  // int
```

**Inspect struct:**
```go
t := reflect.TypeOf(myStruct)
t.NumField()         // Number of fields
t.Field(0).Name      // Field name
t.Field(0).Type      // Field type
t.Field(0).Tag.Get("json")  // Struct tag
```

**Modify value:**
```go
var x int = 10
v := reflect.ValueOf(&x)  // Must use pointer!
v.Elem().SetInt(100)      // x is now 100
```

---

## 2. Unsafe

**Break Go's type safety**

```go
import "unsafe"

// Size of type
unsafe.Sizeof(int(0))     // 8 (on 64-bit)
unsafe.Sizeof(string("")) // 16

// Offset in struct
unsafe.Offsetof(myStruct.field)

// Raw pointer
ptr := unsafe.Pointer(&x)
```

---

## 3. Cgo

**Call C code from Go**

```go
// #include <math.h>
import "C"

func main() {
    result := C.sqrt(16)  // Call C function
}
```

**Requirements:**
- C compiler installed
- Slower than pure Go
- Complex memory management

---

## When to Use?

| Topic | Use When |
|-------|----------|
| Reflect | Building JSON/ORM libraries |
| Unsafe | Extreme performance needs |
| Cgo | Must use C library |

---

## When NOT to Use?

- Normal application code ❌
- When Go alternative exists ❌
- When you're learning ❌

---

## Quick Reference

| Function | Purpose |
|----------|---------|
| `reflect.TypeOf(x)` | Get type |
| `reflect.ValueOf(x)` | Get value |
| `reflect.Kind()` | Base type (int, string, struct) |
| `unsafe.Sizeof(x)` | Size in bytes |
| `unsafe.Pointer(x)` | Raw pointer |
| `import "C"` | Enable cgo |

---

## Summary

```
Reflect = Runtime type inspection
Unsafe  = Break safety rules  
Cgo     = Call C code

All three: USE SPARINGLY! ⚠️
```
