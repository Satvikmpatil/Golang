package q5

import "fmt"

// Run executes Question 5
func Run() {
	fmt.Println("Q5: Positive, Negative, or Zero")

	fmt.Println("Enter the number")

	var a int
	fmt.Scan(&a)

	switch {
	case a > 0:
		fmt.Println("Positive")
	case a < 0:
		fmt.Println("Negative")
	default:
		fmt.Println("Zero")
	}
}