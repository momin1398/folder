//string conversation
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a1, b string
	var c, d float64
	var e, f int
	fmt.Println("Enter string")
	fmt.Scan(&a1)
	c, _ = strconv.ParseFloat(a1, 64)
	fmt.Printf("string to float %#v\n", c)
	fmt.Println("enter float")
	fmt.Scan(&d)
	b = fmt.Sprintf("%v", d)
	fmt.Printf("float to string %#v\n", b)
	e, _ = strconv.Atoi(a1)
	fmt.Printf("string to int %#v\n", e)
	fmt.Println("Enter int")
	fmt.Scan(&f)
	b = strconv.Itoa(f)
	fmt.Printf("int to string %#v\n", b)

}
