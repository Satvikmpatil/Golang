package q4

import "fmt"

// Run executes Question 4
func Run() {
	fmt.Println("Q4: Find Maximum of Three Numbers")
	
	fmt.Println("Enter the number")

	var a int

	fmt.Scan(&a)

	var b int

	fmt.Scan(&b)

	var c int

	fmt.Scan(&c)

	var d int
	if a > b {
		if a > c {
			d = a
		} else {
			d = c
		}
	} else {
		if b > c {
			d = b
		} else {
			d = c
		}
	}
	
	fmt.Println("Maximum:", d)
}
