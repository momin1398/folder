package main

import "fmt"

func main() {
	var num int
	ch := make(chan int)
	fmt.Println("Enter the number:")
	fmt.Scanf("%v", &num)
	for i := 0; i <= num; i++ {
		go func(n int, c chan int) { //anonymous channel

			sum := 0
			for i := 1; i <= n; i++ {
				sum += i

			}
			c <- sum
		}(i, ch)
		fmt.Printf("sum of %v is %v\n", i, <-ch)
	}
}
