package q12

import "fmt"

// Run executes Question 12
func Run() {
	fmt.Println("Q12: Check Palindrome Number")
	fmt.Println("Enter the number")

	var a int
	fmt.Scan(&a)
	
	var i int
	rev :=0
	b := a
	for a !=0 {
		i =a %10;
		rev = rev*10 +i
		a=a/10;
	}
	
	if b == rev{
		fmt.Println("Palindrome")
	}else{
		fmt.Println("Not Palindrome")
	}
	
}
