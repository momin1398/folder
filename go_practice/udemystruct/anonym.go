package main

import "fmt"

func main() {
	name := struct {
		firstname string
		lastname  string
		age       int
		bool
	}{
		firstname: "xxx", lastname: "yyy", age: 23,
	}
	fmt.Printf("The anonymous struct details are %v\n", name)
}
