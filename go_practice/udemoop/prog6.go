package main

import "fmt"

func main() {
	var x interface{}
	x = 6
	fmt.Printf("The interface type & value is %#v\n", x)
	x = 7.9
	fmt.Printf("The interface type & value is %#v\n", x)
	x = []int{6, 8, 9, 10}
	fmt.Printf("The interface type & value is %#v\n", x)
	x = append(x.([]int), 23, 56)

	fmt.Printf("The interface type & value is %#v\n", x)

}
