package q17

import "fmt"

// Run executes Question 17
func Run() {
	fmt.Println("Q17: TODO - Write your code here")
	fmt.Println("Enter the number")

	var a int
	fmt.Scan(&a)

	for i := 1 ; i <= a;i++{
		for j := 0; j< i;j++{
			fmt.Print(j+1)
		}
		fmt.Println()
	}
	// Your code here
	
}
