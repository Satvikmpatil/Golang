package math

// Double returns x * 2 (Exported - Uppercase)
func Double(x int) int {
	return x * 2
}

// Triple returns x * 3 (Exported - Uppercase)
func Triple(x int) int {
	return x * 3
}

// helper is not exported (lowercase)
func helper(x int) int {
	return x + 1
}

func Sum(x int)int{
    return helper(x)
}
