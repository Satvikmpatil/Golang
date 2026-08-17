package q1

import "fmt"

// Run executes Question 1: Hello World
func Run() {
	fmt.Println("Q1: Add Two Numbers")
	fmt.Println("Enter the numbers")
	var a int
	fmt.Scan(&a)
	var b int
	fmt.Scan(&b)
	c:= a+b
	fmt.Println("Sum: ",c)
}
