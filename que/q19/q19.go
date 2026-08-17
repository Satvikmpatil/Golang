package q19

import "fmt"

// Run executes Question 19
func Run() {
	fmt.Println("Q19: Find Second Largest Element")

	fmt.Println("Enter the size")
	var n int
	fmt.Scan(&n)

	a := make([]int, n)

	fmt.Println("Enter the elements")
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	largest := a[0]
	second := a[0]

	for i := 1; i < n; i++ {
		if a[i] > largest {
			second = largest
			largest = a[i]
		} else if a[i] > second && a[i] != largest {
			second = a[i]
		}
	}

	fmt.Println("Second Largest", second)
}