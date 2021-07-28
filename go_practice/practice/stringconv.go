package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 5
	i := strconv.Itoa(a)
	fmt.Printf("The %T to  %T is %v\n", a, i, i) //integr to string
	b := 45.5
	c := fmt.Sprintf("%f", b)
	fmt.Printf("the %T to %T is %v\n", b, c, c)
	k := "42"
	j, _ := strconv.Atoi(k)
	fmt.Printf("the %T to %T is %v\n", k, j, j)
	m := "56.87"
	n, _ := strconv.ParseFloat(m, 64)
	fmt.Printf("the %T to %T is %v\n", m, n, n)

}
