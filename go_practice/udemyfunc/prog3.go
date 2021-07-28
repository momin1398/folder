package main

import (
	"fmt"
	"strconv"
)

func myfunc(n string) int {
	i, _ := strconv.Atoi(n)
	ii, _ := strconv.Atoi(n + n)
	iii, _ := strconv.Atoi(n + n + n)
	return i + ii + iii
}

func main() {
	var a string
	fmt.Println("Enter the string in ")
	fmt.Scan(&a)
	fmt.Printf("The result is %v", myfunc(a))

}
