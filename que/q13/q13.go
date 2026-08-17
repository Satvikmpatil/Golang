package q13

import "fmt"

func Prime(a int)bool{
	if(a<=2){
		return true
	}
	for i:=2;i<a/2;i++{
		if a %i ==0{
			return false
		}
	}
	return true
}

// Run executes Question 13
func Run() {
	fmt.Println("Q13: Check Prime Number")
	fmt.Println("Enter the number")

	var a int
	fmt.Scan(&a)
	fmt.Println(Prime(a))
	// Your code here
	
}
