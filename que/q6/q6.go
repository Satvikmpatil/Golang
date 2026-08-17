package q6

import "fmt"

// Run executes Question 6
func Run() {
	fmt.Println("Q6: Print Numbers 1 to N")

	fmt.Println("Enter the number")

	var a int
	fmt.Scan(&a)

	for i := range a{
		fmt.Println(i+1)
	}
	
	// Your code here
	
}
