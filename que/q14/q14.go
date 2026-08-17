package q14

import "fmt"

func Prime(n int){
	i := 0
	j := 1
	var c int
	for range n{
		fmt.Println(i)
		c = i +j
		i=j
		j=c
	}
}

// Run executes Question 14
func Run() {
	fmt.Println("Q14: Print Fibonacci Series")
	fmt.Println("Enter the number")

	var a int
	fmt.Scan(&a)
	Prime(a)
	
	// Your code here
	
}
