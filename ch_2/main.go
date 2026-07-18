package main

import "fmt"

// ============================================
// CHAPTER 2: Predeclared Types and Declarations
// ============================================

// ----- CONSTANTS (values that never change) -----
const pi = 3.14159       // untyped - flexible, works with any numeric type
const maxUsers int = 100 // typed - only works with int

const (
	appName    = "MyApp"
	appVersion = "1.0.0"
	statusOK   = 200
)

func main() {

	// =====================================
	// 1. ZERO VALUES (default when not assigned)
	// =====================================
	var myInt int       // zero value = 0
	var myFloat float64 // zero value = 0.0
	var myString string // zero value = "" (empty)
	var myBool bool     // zero value = false

	fmt.Println("--- Zero Values ---")
	fmt.Println("int:", myInt)
	fmt.Println("float64:", myFloat)
	fmt.Println("string:", myString)
	fmt.Println("bool:", myBool)

	// =====================================
	// 2. LITERALS (values written directly)
	// =====================================
	fmt.Println("\n--- Literals ---")

	// Integer literals
	decimal := 42       // base 10 (normal)
	binary := 0b1010    // base 2 (binary) = 10
	octal := 0o77       // base 8 (octal) = 63
	hex := 0xFF         // base 16 (hex) = 255
	bigNum := 1_000_000 // underscores for readability

	fmt.Println("decimal:", decimal)
	fmt.Println("binary 0b1010:", binary)
	fmt.Println("octal 0o77:", octal)
	fmt.Println("hex 0xFF:", hex)
	fmt.Println("big number:", bigNum)

	// Float literals
	price := 19.99
	scientific := 6.03e23 // 6.03 × 10^23

	fmt.Println("price:", price)
	fmt.Println("scientific:", scientific)

	// Rune literals (single character - use single quotes)
	letter := 'A'
	emoji := '😀'
	newline := '\n' // special characters

	fmt.Println("letter:", letter, "=", string(letter))
	fmt.Println("emoji:", emoji, "=", string(emoji))
	fmt.Println("newline code:", newline)

	// String literals (text - use double quotes)
	greeting := "Hello, World!"
	withEscape := "Line1\nLine2" // \n = newline
	raw := `Raw string
can have multiple lines
and "quotes" without escaping`

	fmt.Println("greeting:", greeting)
	fmt.Println("withEscape:", withEscape)
	fmt.Println("raw:", raw)

	// =====================================
	// 3. BOOL TYPE (true or false)
	// =====================================
	fmt.Println("\n--- Bool ---")

	isActive := true
	isComplete := false

	fmt.Println("isActive:", isActive)
	fmt.Println("isComplete:", isComplete)

	// =====================================
	// 4. INTEGER TYPES
	// =====================================
	fmt.Println("\n--- Integers ---")

	// Just use 'int' most of the time!
	var age int = 25
	var count int = 100

	// Use specific types only when needed
	var smallNum int8 = 127    // -128 to 127
	var bigNum64 int64 = 90000 // very large numbers
	var positive uint = 50     // unsigned (no negatives)

	// byte = uint8 (for raw data)
	var b byte = 255

	fmt.Println("age:", age)
	fmt.Println("count:", count)
	fmt.Println("smallNum (int8):", smallNum)
	fmt.Println("bigNum64 (int64):", bigNum64)
	fmt.Println("positive (uint):", positive)
	fmt.Println("byte:", b)

	// =====================================
	// 5. FLOAT TYPES
	// =====================================
	fmt.Println("\n--- Floats ---")

	// Always use float64!
	var temperature float64 = 98.6
	var piValue float64 = 3.14159

	fmt.Println("temperature:", temperature)
	fmt.Println("pi:", piValue)

	// WARNING: floats are not exact!
	result := 0.1 + 0.2
	fmt.Println("0.1 + 0.2 =", result) // not exactly 0.3!

	// =====================================
	// 6. STRINGS AND RUNES
	// =====================================
	fmt.Println("\n--- Strings and Runes ---")

	// string = text (double quotes)
	name := "Satvik"

	// rune = one character (single quotes)
	initial := 'S'

	// Strings are immutable (can't change individual chars)
	// name[0] = 'X'  // ERROR! Can't do this

	// But you can reassign
	name = "Go"

	fmt.Println("name:", name)
	fmt.Println("initial:", initial, "=", string(initial))

	// String concatenation
	fullName := "Hello" + " " + "World"
	fmt.Println("fullName:", fullName)

	// =====================================
	// 7. TYPE CONVERSION (explicit only!)
	// =====================================
	fmt.Println("\n--- Type Conversion ---")

	// Go does NOT auto-convert! You must be explicit.
	var x int = 10
	var y float64 = 30.5

	// Can't do: var sum = x + y  // ERROR!
	// Must convert:
	sum1 := float64(x) + y // convert int to float64
	sum2 := x + int(y)     // convert float64 to int (loses decimal!)

	fmt.Println("float64(x) + y =", sum1)
	fmt.Println("x + int(y) =", sum2) // 10 + 30 = 40 (not 40.5!)

	// Even same family needs conversion
	var a int = 10
	var b64 int64 = 20
	sum3 := int64(a) + b64 // must convert int to int64

	fmt.Println("int64(a) + b64 =", sum3)

	// =====================================
	// 8. VAR vs := (declaring variables)
	// =====================================
	fmt.Println("\n--- var vs := ---")

	// Use 'var' for zero values
	var counter int // clearly want zero

	// Use ':=' for values (inside functions)
	message := "Hello" // short and clean
	number := 42
	active := true

	// Use 'var' when you need specific type
	var specificByte byte = 25 // want byte, not int

	fmt.Println("counter:", counter)
	fmt.Println("message:", message)
	fmt.Println("number:", number)
	fmt.Println("active:", active)
	fmt.Println("specificByte:", specificByte)

	// Multiple variables
	a1, b1 := 10, "hello"
	fmt.Println("a1:", a1, "b1:", b1)

	// =====================================
	// 9. CONSTANTS
	// =====================================
	fmt.Println("\n--- Constants ---")

	// Constants never change
	const localPi = 3.14159
	const maxSize = 1000

	// Untyped const is flexible
	const flexNum = 10
	var asInt int = flexNum       // works!
	var asFloat float64 = flexNum // works too!

	fmt.Println("localPi:", localPi)
	fmt.Println("maxSize:", maxSize)
	fmt.Println("flexNum as int:", asInt)
	fmt.Println("flexNum as float:", asFloat)

	// Using package-level constants
	fmt.Println("appName:", appName)
	fmt.Println("maxUsers:", maxUsers)

	// =====================================
	// 10. OPERATORS
	// =====================================
	fmt.Println("\n--- Operators ---")

	// Math operators
	fmt.Println("10 + 3 =", 10+3)
	fmt.Println("10 - 3 =", 10-3)
	fmt.Println("10 * 3 =", 10*3)
	fmt.Println("10 / 3 =", 10/3) // integer division = 3
	fmt.Println("10 % 3 =", 10%3) // remainder = 1

	// Comparison operators
	fmt.Println("5 == 5:", 5 == 5)
	fmt.Println("5 != 3:", 5 != 3)
	fmt.Println("5 > 3:", 5 > 3)
	fmt.Println("5 < 3:", 5 < 3)

	// Shorthand operators
	num := 10
	num += 5 // num = num + 5
	fmt.Println("num += 5:", num)
	num *= 2 // num = num * 2
	fmt.Println("num *= 2:", num)

	// =====================================
	// 11. BYTE vs RUNE
	// =====================================
	fmt.Println("\n--- Byte vs Rune ---")

	// byte = small (1 byte) - ASCII only
	var asciChar byte = 'A' // works for English

	// rune = big (4 bytes) - any character
	var anyChar rune = '語' // works for any language
	var emojiChar rune = '🎉'

	fmt.Println("byte 'A':", asciChar)
	fmt.Println("rune '語':", anyChar, "=", string(anyChar))
	fmt.Println("rune '🎉':", emojiChar, "=", string(emojiChar))

	// =====================================
	// SUMMARY
	// =====================================
	fmt.Println("\n--- Quick Reference ---")
	fmt.Println("int    - whole numbers (use this!)")
	fmt.Println("float64 - decimals (use this!)")
	fmt.Println("string - text")
	fmt.Println("bool   - true/false")
	fmt.Println("byte   - small char/data (ASCII)")
	fmt.Println("rune   - any character (Unicode)")
	fmt.Println(":=     - declare with value")
	fmt.Println("var    - declare with zero value or specific type")
	fmt.Println("const  - never changes")
}
