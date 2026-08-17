package q7

import "fmt"

// Run executes Question 7
func Run() {
	fmt.Println("Q7: Sum of First N Natural Numbers")

	fmt.Println("Enter the number")

	var a int
	fmt.Scan(&a)

	sum := (a*(a+1))/2

	fmt.Println(sum)
	
	// Your code here
	
}
