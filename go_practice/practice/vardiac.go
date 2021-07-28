package main

import (
	"fmt"
	"strings"
)

func sumpr(n ...int) (int, int) {
	sum := 0
	pr := 1
	for _, v := range n {
		sum += v
		pr *= v
	}
	return sum, pr

}
func varcom(age int, n ...string) string {
	d := strings.Join(n, " ")
	d = strings.ToUpper(d)
	returnString := fmt.Sprintf("The age is %v of %v", age, d)
	return returnString

}

func main() {
	num := []int{6, 8, 9, 34}
	num = append(num, 76, 98, 3, 1)
	fmt.Printf("The array is %v\n", num)
	sum, prod := sumpr(num...)
	fmt.Printf("The sum is %v,product is %v\n", sum, prod)
	name := []string{"xxx", "yyyy", "bbbb"}
	a := 23
	s := varcom(a, name...)
	fmt.Println(s)

}
