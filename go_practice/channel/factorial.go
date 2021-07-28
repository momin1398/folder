package main

import (
	"fmt"
	"strings"
)

func factorial(n int, c chan int) {
	f := 1
	for i := 1; i <= n; i++ {
		f *= i
	}
	c <- f

}

func main() {
	ch := make(chan int)
	defer close(ch)
	go factorial(5, ch)
	fmt.Printf("The factorial of 5 is: %v\n", <-ch)
	for i := 1; i <= 10; i++ {
		go factorial(i, ch)
		fmt.Printf("The factorial of %v is: %v\n", i, <-ch)
	}
	fmt.Println(strings.Repeat("#", 20))

	for i := 10; i >= 1; i-- {
		go func(n int, c chan int) {
			f := 1
			for i := n; i >= 1; i-- {
				f *= i
			}
			c <- f

		}(i, ch)
		fmt.Printf("The factorial of %v is %v \n", i, <-ch)
	}
}
