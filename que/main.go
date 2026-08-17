package main

import (
	"fmt"

	"que/q1"
	"que/q10"
	"que/q11"
	"que/q12"
	"que/q13"
	"que/q14"
	"que/q15"
	"que/q16"
	"que/q17"
	"que/q18"
	"que/q19"
	"que/q2"
	"que/q20"
	"que/q21"
	"que/q22"
	"que/q23"
	"que/q24"
	"que/q25"
	"que/q26"
	"que/q27"
	"que/q28"
	"que/q29"
	"que/q3"
	"que/q30"
	"que/q31"
	"que/q32"
	"que/q33"
	"que/q34"
	"que/q35"
	"que/q36"
	"que/q37"
	"que/q38"
	"que/q39"
	"que/q4"
	"que/q40"
	"que/q41"
	"que/q42"
	"que/q43"
	"que/q44"
	"que/q45"
	"que/q46"
	"que/q47"
	"que/q48"
	"que/q49"
	"que/q5"
	"que/q50"
	"que/q6"
	"que/q7"
	"que/q8"
	"que/q9"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("       50 Fundamental Go Questions      ")
	fmt.Println("========================================")
	fmt.Print("Select Question (1-50): ")

	var choice int
	fmt.Scan(&choice)
	fmt.Println()

	switch choice {
	case 1:
		q1.Run()
	case 2:
		q2.Run()
	case 3:
		q3.Run()
	case 4:
		q4.Run()
	case 5:
		q5.Run()
	case 6:
		q6.Run()
	case 7:
		q7.Run()
	case 8:
		q8.Run()
	case 9:
		q9.Run()
	case 10:
		q10.Run()
	case 11:
		q11.Run()
	case 12:
		q12.Run()
	case 13:
		q13.Run()
	case 14:
		q14.Run()
	case 15:
		q15.Run()
	case 16:
		q16.Run()
	case 17:
		q17.Run()
	case 18:
		q18.Run()
	case 19:
		q19.Run()
	case 20:
		q20.Run()
	case 21:
		q21.Run()
	case 22:
		q22.Run()
	case 23:
		q23.Run()
	case 24:
		q24.Run()
	case 25:
		q25.Run()
	case 26:
		q26.Run()
	case 27:
		q27.Run()
	case 28:
		q28.Run()
	case 29:
		q29.Run()
	case 30:
		q30.Run()
	case 31:
		q31.Run()
	case 32:
		q32.Run()
	case 33:
		q33.Run()
	case 34:
		q34.Run()
	case 35:
		q35.Run()
	case 36:
		q36.Run()
	case 37:
		q37.Run()
	case 38:
		q38.Run()
	case 39:
		q39.Run()
	case 40:
		q40.Run()
	case 41:
		q41.Run()
	case 42:
		q42.Run()
	case 43:
		q43.Run()
	case 44:
		q44.Run()
	case 45:
		q45.Run()
	case 46:
		q46.Run()
	case 47:
		q47.Run()
	case 48:
		q48.Run()
	case 49:
		q49.Run()
	case 50:
		q50.Run()
	default:
		fmt.Println("Invalid choice! Enter 1-50")
	}
}
