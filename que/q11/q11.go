package q11

import "fmt"

// Run executes Question 11
func Run() {
	fmt.Println("Q11: Reverse a Number")

	fmt.Println("Enter the number")

	var a int
	fmt.Scan(&a)
	var i int
	rev :=0

	for a !=0 {
		i =a %10;
		rev = rev*10 +i
		a=a/10;
	}
	fmt.Println(rev)
	// Your code here
	
}
