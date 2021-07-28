package main

import "fmt"

var n = [10]int{}

func init() {
	fmt.Println("The init function excuting")
	for i := 0; i < 10; i++ {
		n[i] = i * 2
	}

}
func main() {
	fmt.Println("The value of array is", n)
}
