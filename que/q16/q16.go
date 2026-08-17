package q16

import "fmt"

// Run executes Question 16
func Run() {
	fmt.Println("Q16: Right Triangle Star Pattern")

	fmt.Println("Enter the number")

	var a int
	fmt.Scan(&a)
	
	for i := 1 ; i <= a;i++{
		for j := 0; j< i;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
	
	// Your code here
	
}
