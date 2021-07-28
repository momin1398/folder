package main

import "fmt"

func swap(a, b *float64) {
	*a, *b = *b, *a
}

func main() {
	//code1
	x := 10.10
	ptr := &x
	fmt.Printf("The type is %#v , value is %v in ptr \n", ptr, *ptr)
	fmt.Println("The address of ptr", &ptr)
	//code2
	n, y := 10, 2
	ptr1, ptry := &n, &y
	z := *ptr1 / *ptry
	fmt.Printf("z is %d\n", z)
	//code3
	a, b := 5.5, 8.8
	fmt.Println("before swapping", a, b)
	swap(&a, &b)
	fmt.Println("After swapping", a, b)
}
