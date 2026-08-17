package q10

import "fmt"

// Run executes Question 10
func Run() {
	fmt.Println("Q10: Count Digits in a Number")
	
	// Your code here
	
	fmt.Println("Enter the number")

	var a int
	fmt.Scan(&a)
	i :=0
	for a != 0{
		a = a/10
		i = i+1
	}
	fmt.Println(i)
}
