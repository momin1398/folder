package main

import (
	"fmt"
)

var a []int

func push() {

	fmt.Println("Enter num of elements")
	var j, n int
	fmt.Scan(&n)
	fmt.Println("Enter array")
	for i := 0; i < n; i++ {
		fmt.Scan(&j)
		a = append(a, j)
	}

	fmt.Printf("The  array is %v\n", a)

}
func pop() {
	fmt.Println("Enter the element to be del")
	var j, m int
	fmt.Scan(&j)
	if len(a) > 0 {

		m = len(a)
		fmt.Printf("the removed value are %v\n", a[m-j:])

		a = a[:m-j]
		fmt.Printf("The new array is %v\n", a)
	}

}
func main() {

	push()
	pop()

}
