package q3

import "fmt"

// Run executes Question 3
func Run() {
	fmt.Println("Q3: Check Even or Odd Using")

	fmt.Println("Enter the number")

	var a int

	fmt.Scan(&a)

	if a % 2 == 0 {
		fmt.Println("Even")
	} else{
		fmt.Println("Odd")
	}
	
	// Your code here
	
}
