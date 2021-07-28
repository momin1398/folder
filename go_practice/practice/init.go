package main

import "fmt"

var arr [10]int

func init() {
	fmt.Println("Enter array")
	for i := 0; i < 10; i++ {
		fmt.Scan(&arr[i])
	}
}
func add(a [10]int) int {
	s := 0
	for _, i := range a {
		s += i
	}
	return s

}
func addv(a ...int) int {
	s := 0
	for _, i := range a {
		s += i
	}
	return s

}
func main() {
	fmt.Println("The array are %v", arr)
	s := add(arr)
	fmt.Printf("The sum is %v\n", s)
	var a []int
	a = append(a, arr[:]...)
	a = append(a, 7, 8, 9)
	fmt.Printf("calling vardiac func \n")
	sum := addv(a...)
	fmt.Printf("The sum of slice %v\n", sum)

}
