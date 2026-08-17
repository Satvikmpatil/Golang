package q2

import "fmt"

// Run executes Question 2: Add Two Numbers
func Run() {
	fmt.Println("Q2: Swap")
	fmt.Println()

	var a, b int
	fmt.Print("Enter first number: ")
	fmt.Scan(&a)
	fmt.Print("Enter second number: ")
	fmt.Scan(&b)

	c := a + b
	a = c - a
	b = c - a
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
