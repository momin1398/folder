package main

import "fmt"

func cube(a int) int {
	return a * a * a
}
func factsum(a int) (int, int) {
	fa := 1
	m := 0
	for i := 1; i <= a; i++ {
		fa *= i
		m += i
	}
	return fa, m
}

func main() {
	var a int
	println("Enter the num")
	fmt.Scan(&a)
	//code 1
	fmt.Printf("The cube of %v is %v\n", a, cube(a))
	//code2
	fact, sum := factsum(a)
	fmt.Printf("The factorial of %v is %v\n", a, fact)
	fmt.Printf("The sum of %v is %v\n", a, sum)

}
