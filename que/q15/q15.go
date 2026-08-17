package q15

import "fmt"

// Run executes Question 15
func Run() {
	fmt.Println("Q15: Print Fibonacci Series")
	var a, b int
	fmt.Print("Enter first number: ")
	fmt.Scan(&a)
	fmt.Print("Enter second number: ")
	fmt.Scan(&b)
	// Your code here
	for b != 0 {
		r := a % b
		a = b
		b = r
	}
	fmt.Println("GCD: ", a)
	
}
