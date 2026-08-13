package main

import "fmt"

// ============================================
// CHAPTER 7: Types, Methods, and Interfaces
// ============================================

// =====================================
// 1. USER-DEFINED TYPES
// =====================================

type Score int                    // Score is based on int
type Converter func(string) int   // Function type
type TeamScores map[string]int    // Map type

// =====================================
// 2. STRUCT TYPE
// =====================================

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// =====================================
// 3. METHODS - Value Receiver vs Pointer Receiver
// =====================================

// Value receiver - doesn't modify original
func (p Person) FullName() string {
	return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}

// Pointer receiver - CAN modify original
func (p *Person) SetAge(age int) {
	p.Age = age
}

// String method (implements fmt.Stringer interface)
func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

// =====================================
// 4. COUNTER EXAMPLE - Pointer vs Value Receiver
// =====================================

type Counter struct {
	total int
}

// Pointer receiver - modifies counter
func (c *Counter) Increment() {
	c.total++
}

// Value receiver - doesn't modify (useless for increment!)
func (c Counter) IncrementWrong() {
	c.total++ // Only changes copy!
}

func (c Counter) Total() int {
	return c.total
}

// =====================================
// 5. NIL RECEIVER - Binary Tree Example
// =====================================

type IntTree struct {
	val         int
	left, right *IntTree
}

// Can handle nil receiver!
func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}
	return it
}

func (it *IntTree) Contains(val int) bool {
	if it == nil {
		return false
	}
	if val < it.val {
		return it.left.Contains(val)
	}
	if val > it.val {
		return it.right.Contains(val)
	}
	return true
}

// =====================================
// 6. METHOD AS VALUE
// =====================================

type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}

// =====================================
// 7. TYPE BASED ON TYPE (NOT Inheritance!)
// =====================================

type HighScore Score // HighScore based on Score, but different type!

// =====================================
// 8. IOTA - For Enumerations
// =====================================

type Day int

const (
	_         Day = iota // 0 (skip)
	Monday               // 1
	Tuesday              // 2
	Wednesday            // 3
	Thursday             // 4
	Friday               // 5
	Saturday             // 6
	Sunday               // 7
)

// =====================================
// 9. EMBEDDING - Composition
// =====================================

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee        // Embedded! No field name
	Reports  []Employee
}

// =====================================
// 10. INTERFACES
// =====================================

// Interface = list of methods
type Speaker interface {
	Speak()
}

// Dog implements Speaker (has Speak method)
type Dog struct {
	Name string
}

func (d Dog) Speak() {
	fmt.Println(d.Name, "says: Woof!")
}

// Cat implements Speaker (has Speak method)
type Cat struct {
	Name string
}

func (c Cat) Speak() {
	fmt.Println(c.Name, "says: Meow!")
}

// Function accepts interface - works with any Speaker!
func MakeSound(s Speaker) {
	s.Speak()
}

// =====================================
// 11. STRINGER INTERFACE (from fmt package)
// =====================================

type Product struct {
	Name  string
	Price float64
}

// Implements fmt.Stringer interface
func (p Product) String() string {
	return fmt.Sprintf("%s: $%.2f", p.Name, p.Price)
}

// =====================================
// MAIN FUNCTION
// =====================================

func main() {

	// --- 1. User-Defined Types ---
	fmt.Println("--- 1. User-Defined Types ---")
	var score Score = 100
	fmt.Println("Score:", score)

	// --- 2. Struct and Methods ---
	fmt.Println("\n--- 2. Struct and Methods ---")
	person := Person{FirstName: "Bob", LastName: "Smith", Age: 30}
	fmt.Println("FullName:", person.FullName())
	fmt.Println("String:", person.String())

	// --- 3. Pointer vs Value Receiver ---
	fmt.Println("\n--- 3. Pointer vs Value Receiver ---")
	person.SetAge(35) // Modifies original!
	fmt.Println("After SetAge(35):", person.Age)

	// --- 4. Counter Example ---
	fmt.Println("\n--- 4. Counter Example ---")
	var c Counter

	c.IncrementWrong()
	fmt.Println("After IncrementWrong:", c.Total()) // Still 0!

	c.Increment()
	fmt.Println("After Increment:", c.Total()) // 1!

	c.Increment()
	fmt.Println("After Increment:", c.Total()) // 2!

	// --- 5. nil Receiver (Binary Tree) ---
	fmt.Println("\n--- 5. nil Receiver (Binary Tree) ---")
	var tree *IntTree // nil!
	tree = tree.Insert(5)
	tree = tree.Insert(3)
	tree = tree.Insert(7)
	fmt.Println("Contains 3:", tree.Contains(3)) // true
	fmt.Println("Contains 9:", tree.Contains(9)) // false

	// --- 6. Method as Value ---
	fmt.Println("\n--- 6. Method as Value ---")
	myAdder := Adder{start: 10}
	addFunc := myAdder.AddTo // Save method to variable
	fmt.Println("addFunc(5):", addFunc(5))   // 15
	fmt.Println("addFunc(10):", addFunc(10)) // 20

	// --- 7. Type Based on Type ---
	fmt.Println("\n--- 7. Type Based on Type ---")
	var s Score = 100
	var hs HighScore = 200
	// hs = s        // Error! Different types
	hs = HighScore(s) // OK with conversion
	fmt.Println("HighScore:", hs)

	// --- 8. iota ---
	fmt.Println("\n--- 8. iota ---")
	fmt.Println("Monday:", Monday)       // 1
	fmt.Println("Friday:", Friday)       // 5
	fmt.Println("Sunday:", Sunday)       // 7

	// --- 9. Embedding ---
	fmt.Println("\n--- 9. Embedding ---")
	m := Manager{
		Employee: Employee{
			Name: "Alice",
			ID:   "M001",
		},
		Reports: []Employee{},
	}
	fmt.Println("Name:", m.Name)               // Direct access!
	fmt.Println("ID:", m.ID)                   // Direct access!
	fmt.Println("Description:", m.Description()) // Method works!

	// --- 10. Interfaces ---
	fmt.Println("\n--- 10. Interfaces ---")
	dog := Dog{Name: "Buddy"}
	cat := Cat{Name: "Whiskers"}

	MakeSound(dog) // Works!
	MakeSound(cat) // Works!

	// --- 11. Stringer Interface ---
	fmt.Println("\n--- 11. Stringer Interface ---")
	product := Product{Name: "Laptop", Price: 999.99}
	fmt.Println(product) // Uses String() method automatically!

	// --- Summary ---
	fmt.Println("\n--- Summary ---")
	fmt.Println("Type = name for kind of data")
	fmt.Println("Method = function attached to type")
	fmt.Println("Value receiver = doesn't modify")
	fmt.Println("Pointer receiver = can modify")
	fmt.Println("Interface = list of methods")
	fmt.Println("Embedding = composition (not inheritance)")
}
