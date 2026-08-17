package q18

import "fmt"

// Run executes Question 18
func Run() {
	fmt.Println("Q18: Find Largest Element in Slice")

	fmt.Println("Enter the size:")
	var n int
	fmt.Scan(&n)

	a := make([]int, n)

	fmt.Println("Enter the elements:")
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	m := a[0]

	for i := 1; i < n; i++ {
		if a[i] > m {
			m = a[i]
		}
	}

	fmt.Println("MAX", m)
}