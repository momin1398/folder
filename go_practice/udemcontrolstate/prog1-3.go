package main

import "fmt"

func main() {
	//code 1,
	fmt.Println("The num divisible by 7 from 1-50 are ")
	for i := 1; i <= 50; i++ {
		if i%7 == 0 {
			fmt.Printf("%v\n", i)
		}
	}
	//code 2
	fmt.Println("The num divisible by 7 from 1-50 using continue statement are ")
	for i := 1; i <= 50; i++ {
		if i%7 != 0 {
			continue
		}
		fmt.Printf("%v\n", i)
	}
	//code 3
	fmt.Println("The first three num divisible by 7 from 1-50 break statement are ")
	j := 0
	for i := 1; i <= 50; i++ {
		if i%7 == 0 {
			fmt.Printf("%v\n", i)
			j++
		}
		if j == 3 {
			break
		}
	}
}
