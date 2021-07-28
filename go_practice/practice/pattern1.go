package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Enter the num")
	var n int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v ", j)
		}
		fmt.Println()
	}
	fmt.Println(strings.Repeat("#", 25))
	for i := 1; i <= n; i++ {
		for j := n; j >= i; j-- {
			fmt.Printf("%v ", j)
		}
		fmt.Println()

	}
}
