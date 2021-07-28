package main

import "fmt"

func add(a ...int) int {
	s := 0
	for _, i := range a {
		s += i
	}
	return s
}
func prod(a ...int) (s int) {
	s = 1
	for _, i := range a {
		s *= i
	}
	return
}

func main() {
	a := []int{6, 8, 9}
	a = append(a, 7, 5, 8, 5)
	//code4
	fmt.Printf("The sum is %v\n", add(a...))
	//code5
	fmt.Printf("The prod is %v\n", prod(a...))
}
