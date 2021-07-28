package main

import (
	"fmt"
	"strings"
)

func getDetails() [10][10]int {
	var a [10][10]int

	fmt.Println("Enter array")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scan(&a[i][j])

		}
	}
	return a
}
func sum(a, b [10][10]int) {
	fmt.Println("The added array")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%v ", a[i][j]+b[i][j])

		}
		fmt.Println()
	}

}
func subt(a, b [10][10]int) {
	fmt.Println("The subtraction array")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%v ", a[i][j]-b[i][j])

		}
		fmt.Println()
	}

}
func mul(a, b [10][10]int) {
	fmt.Println("The multipied array")
	var m [10][10]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			m[i][j] = 0
			for k := 0; k < 3; k++ {
				m[i][j] += a[i][k] * b[k][j]

			}
			fmt.Printf("%v ", m[i][j])
		}
		fmt.Println()
	}

}

func main() {

	a := getDetails()
	b := getDetails()
	var c string
	fmt.Println("enter")
	fmt.Scan(&c)
	c = strings.ToLower(c)
	switch c {
	case "addition":
		sum(a, b)
	case "subtraction":
		subt(a, b)
	case "multiply":
		mul(a, b)
	default:
		fmt.Println("not found")
	}

}
