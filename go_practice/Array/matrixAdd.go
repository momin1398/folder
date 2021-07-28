package main

import "fmt"

func main() {
	var n, m int
	fmt.Print("Enter the value:")
	fmt.Scan(&n)
	fmt.Print("Enter the value:")
	fmt.Scan(&m)
	var a [10][10]int
	var b [10][10]int
	fmt.Println("Enter the elements of first matrix")
	for i := 0; i < n; i++ {

		for j := 0; j < m; j++ {
			fmt.Scan(&a[i][j])
		}

	}

	fmt.Println("Enter the elements of second matrix")
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&b[i][j])
		}

	}
	fmt.Println("The addition of two matrices is")
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%v ", a[i][j]+b[i][j])
		}
		fmt.Println()

	}

}
