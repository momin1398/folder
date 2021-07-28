package main

import (
	"fmt"
	"strings"
)

func addition(a int, b int) int { //normal function
	return a + b
}
func addprod(a ...int) (int, int) { //vardiac func
	sum := 0
	prod := 1
	for _, i := range a {
		sum += i
		prod *= i

	}
	return sum, prod
}

func varcomb(a int, names ...string) string { //combination of vardiac & non vardiac
	b := strings.Join(names, " ")
	b = strings.ToUpper(b)
	returnString := fmt.Sprintf("THE AGE OF %v IS %v", b, a)
	return returnString
}
func main() {
	var num1 int
	var num2 int
	fmt.Println("Enter two numbers")
	fmt.Scanf("%v %v", &num1, &num2)
	add := addition(num1, num2)
	fmt.Printf("The addition of %v and %v is %v\n", num1, num2, add)
	arr := []int{5, 7, 8}
	arr = append(arr, 8, 9, 10)
	sum, prod := addprod(arr...)
	fmt.Printf("The sum of array is %v\nThe prod of array is %v\n", sum, prod)
	name := []string{"xxxx", "yyyy", "zzzz"}
	age := 23
	fmt.Println(varcomb(age, name...))

}
