package main

import "fmt"

func sf(n int) (int, int) {
	s := 0
	p := 1

	for i := 1; i <= n; i++ {
		s += i
		p *= i

	}
	return s, p
}

func main() {
	var n int
	fmt.Println("Enter the num")
	fmt.Scan(&n)
	s, f := sf(n)
	fmt.Printf("The sum of %v is %v\nThe factorial of %v is %v", n, s, n, f)
}
