package main

import "fmt"

var a [10][10]int
var b [10][10]int
var c [10][10]int

func sub(n int, m int) {
	fmt.Println("The subtraction of 2 array ")
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			c[i][j] = a[i][j] - b[i][j]
			fmt.Printf("%+d ", c[i][j])
		}
		fmt.Println()
	}

}

func main() {
	fmt.Println("Enter the size")
	var n, m int
	fmt.Scan(&n, &m)
	fmt.Println("enter first array")
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&a[i][j])
		}
	}

	fmt.Println("enter second array")
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&b[i][j])
		}
	}

	sub(n, m)

}
