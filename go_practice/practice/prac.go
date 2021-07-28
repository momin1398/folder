package main

import "fmt"

func addpr(a int, b int) (int, int) {
	return a + b, a * b
}
func addPro(c ...int) (int, int) {
	s, p := 0, 1
	for _, i := range c {
		s += i
		p *= i
	}
	return s, p

}
func main() {
	fmt.Println("Enter 2 numbers")
	var a, b int
	fmt.Scan(&a, &b)
	sum, prod := addpr(a, b)
	fmt.Printf("The sum is %v, the prod is %v\n", sum, prod)
	var c []int
	fmt.Println("Enter array")
	var m int
	for i := 0; i < 5; i++ {
		fmt.Scan(&m)
		c = append(c, m)

	}
	s, p := addPro(c...)
	fmt.Printf("The sum of array %v,the prod of array %v\n", s, p)

}
