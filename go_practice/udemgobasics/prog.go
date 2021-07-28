package main

import "fmt"

func main() {
	x := 10
	y := 15.5
	z := "Xyz"
	fmt.Println(x, y, z)
	//code 2
	a, b, c := x, y, z
	_, _, _ = a, b, c
}
