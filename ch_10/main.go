package main

import (
	"fmt"

	"ch_10/math"   // Import local package
	"ch_10/utils"  // Import another local package
)

func main() {
	// --- Using math package ---
	fmt.Println("--- Math Package ---")
	result := math.Double(5)
	fmt.Println("Double(5):", result)

	result = math.Triple(5)
	fmt.Println("Triple(5):", result)

	// --- Using utils package ---
	fmt.Println("\n--- Utils Package ---")
	fmt.Println("Greet:", utils.Greet("Satvik"))
	fmt.Println("Max:", utils.Max(10, 20))

	// --- Summary ---
	fmt.Println("\n--- Summary ---")
	fmt.Println("Uppercase = Exported (public)")
	fmt.Println("lowercase = Not exported (private)")
	fmt.Println("Import: module/package")
	fmt.Println("Use: package.Function()")
}
