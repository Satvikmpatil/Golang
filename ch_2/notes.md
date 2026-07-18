# Chapter 2: Predeclared Types and Declarations

## Zero Values

When you declare a variable without a value, Go gives it a **default value**.

| Type | Zero Value |
|------|------------|
| `int` | `0` |
| `float64` | `0.0` |
| `string` | `""` (empty) |
| `bool` | `false` |

```go
var x int       // x = 0
var s string    // s = ""
var b bool      // b = false
```

---

## Literals

**Literal = Value written directly in code**

### Integer Literals
```go
42          // decimal (base 10)
0b1010      // binary (base 2) = 10
0o77        // octal (base 8) = 63
0xFF        // hex (base 16) = 255
1_000_000   // underscores for readability
```

### Float Literals
```go
3.14        // decimal
6.03e23     // scientific (6.03 × 10^23)
```

### Rune Literals (single character)
```go
'A'         // letter
'😀'        // emoji
'\n'        // newline
'\t'        // tab
```

### String Literals
```go
"Hello"              // interpreted string
`Raw string`         // raw string (can have newlines)
```

---

## Bool Type

```go
var isActive bool = true
var isDone bool = false
```

Only two values: `true` or `false`

---

## Numeric Types

### Integers - Just use `int`!

| Type | Size | Range |
|------|------|-------|
| `int8` | 1 byte | -128 to 127 |
| `int16` | 2 bytes | -32K to 32K |
| `int32` | 4 bytes | -2B to 2B |
| `int64` | 8 bytes | very big |
| `int` | depends | **Use this!** |

### Unsigned (no negatives)
| Type | Size | Range |
|------|------|-------|
| `uint8` / `byte` | 1 byte | 0 to 255 |
| `uint16` | 2 bytes | 0 to 65K |
| `uint32` | 4 bytes | 0 to 4B |
| `uint64` | 8 bytes | very big |

### Floats - Just use `float64`!

| Type | Use |
|------|-----|
| `float32` | Rarely |
| `float64` | **Always use this!** |

**Warning:** Floats are NOT exact! Don't use for money.

---

## Strings and Runes

| Type | Quotes | Example | Use For |
|------|--------|---------|---------|
| `string` | Double `" "` | `"Hello"` | Text |
| `rune` | Single `' '` | `'A'` | One character |

### Strings are immutable
```go
s := "hello"
s = "world"     // OK - new string
s[0] = 'W'      // ERROR - can't change chars
```

### Concatenation
```go
full := "Hello" + " " + "World"
```

---

## Byte vs Rune

| Type | Size | Use For |
|------|------|---------|
| `byte` | 1 byte | ASCII only (English) |
| `rune` | 4 bytes | Any character (emoji, Chinese) |

```go
var a byte = 'A'      // OK
var b byte = '語'     // ERROR - too big!
var c rune = '語'     // OK - rune handles all
```

---

## Type Conversion

**Go does NOT auto-convert! You must be explicit.**

```go
var x int = 10
var y float64 = 20.5

// ERROR: sum := x + y

// Correct:
sum1 := float64(x) + y    // convert int to float
sum2 := x + int(y)        // convert float to int (loses .5!)
```

### Syntax
```go
float64(x)    // convert to float64
int(y)        // convert to int
byte(z)       // convert to byte
string(r)     // convert rune to string
```

---

## var vs :=

### Use `var` when:
- Zero value intended: `var count int`
- Need specific type: `var x byte = 25`
- Outside functions (required)

### Use `:=` when:
- Inside functions with a value: `x := 10`

```go
// Zero value - use var
var count int           // count = 0

// With value - use :=
name := "Go"            // short and clear

// Specific type - use var
var age byte = 25       // need byte, not int
```

### Multiple variables
```go
var a, b int = 10, 20
x, y := 10, "hello"
```

---

## Constants

**`const` = value that NEVER changes**

```go
const pi = 3.14159
const maxSize = 100

pi = 3.0    // ERROR! Can't change const
```

### Typed vs Untyped

```go
// Untyped - flexible
const x = 10
var a int = x       // OK
var b float64 = x   // OK

// Typed - strict
const y int = 10
var c int = y       // OK
var d float64 = y   // ERROR!
```

---

## Operators

### Math
| Operator | Meaning | Example |
|----------|---------|---------|
| `+` | Add | `5 + 3 = 8` |
| `-` | Subtract | `5 - 3 = 2` |
| `*` | Multiply | `5 * 3 = 15` |
| `/` | Divide | `5 / 3 = 1` |
| `%` | Remainder | `5 % 3 = 2` |

### Comparison
| Operator | Meaning |
|----------|---------|
| `==` | Equal |
| `!=` | Not equal |
| `>` | Greater |
| `<` | Less |
| `>=` | Greater or equal |
| `<=` | Less or equal |

### Shorthand
```go
x := 10
x += 5    // x = 15
x -= 3    // x = 12
x *= 2    // x = 24
x /= 4    // x = 6
```

---

## Naming Rules

### Style
- Use `camelCase`: `userName`, `maxRetries`
- NOT `snake_case`: ~~`user_name`~~
- NOT `ALL_CAPS`: ~~`MAX_SIZE`~~

### Short names for small scope
```go
for i := 0; i < 10; i++ { }   // i is fine
```

### Descriptive names for large scope
```go
var userCount int             // package level
var maxConnections int
```

---

## Unused Variables

**Go won't compile if you declare but don't use a variable!**

```go
x := 10          // ERROR if x is never used
```

**But unused constants are OK.**

---

## Quick Reference

| What | Use |
|------|-----|
| Whole numbers | `int` |
| Decimals | `float64` |
| Text | `string` |
| One character | `rune` |
| True/false | `bool` |
| Raw data | `byte` |
| Fixed value | `const` |
| With value (in function) | `:=` |
| Zero value | `var x int` |
| Specific type | `var x byte = 10` |

---

## Key Takeaways

1. **Use `int` for integers** - don't overthink, just use `int`

2. **Use `float64` for decimals** - more precision, less bugs

3. **Zero values are automatic** - `int` = 0, `string` = "", `bool` = false

4. **Go needs explicit conversion** - `float64(x)` not automatic

5. **`:=` inside functions, `var` for zero values**
   ```go
   x := 10         // has value → :=
   var count int   // want zero → var
   ```

6. **Single quotes ≠ Double quotes**
   - `'A'` = rune (one character)
   - `"A"` = string (text)

7. **`byte` for ASCII, `rune` for everything**
   - English only? `byte`
   - Emoji/Chinese? `rune`

8. **Strings are immutable** - can't change `s[0]`, must create new string

9. **`const` is compile-time only** - can't use runtime values

10. **Must use declared variables** - or code won't compile

11. **Use camelCase** - `userName` not `user_name`

12. **Floats are NOT exact** - never use for money!
