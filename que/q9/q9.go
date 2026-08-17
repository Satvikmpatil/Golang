package q9

import "fmt"


func fac(a int)int{
	if a <= 1{
		return 1
	}
	f := 1
	for i := 1; i<=a; i++{
		f = f*i
	}
	return f
}
// Run executes Question 9
func Run() {
	fmt.Println("Q9: Factorial of a Number")

	fmt.Println("Enter the number")

	var a int
	fmt.Scan(&a)
	fmt.Println(fac(a))
	// Your code here
	
}
