package q20

import "fmt"

// Run executes Question 20
func Run() {
	fmt.Println("Q20: Reverse a Slice")

	fmt.Println("Enter the size:")
	var n int
	fmt.Scan(&n)

	a := make([]int, n)

	fmt.Println("Enter the elements:")
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	i := 0
	j := n - 1

	for i < j {
		temp := a[i]
		a[i] = a[j]
		a[j] = temp

		i++
		j--
	}

	fmt.Println("Reversed:", a)
}