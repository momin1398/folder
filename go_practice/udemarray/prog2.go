package main

import "fmt"

func main() {
	//number of positive even num
	nums := [...]int{30, -1, -6, 90, -6}
	s := 0
	for _, j := range nums {
		if j > 0 && j%2 == 0 {
			s++
		}
	}
	fmt.Println("The number of positive even num is ", s)
}
