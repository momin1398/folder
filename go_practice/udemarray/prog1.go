package main

import "fmt"

func main() {
	var cities [2]string
	fmt.Printf("%v\n", cities)
	grade := [3]float64{7.6, 8.9, 9}
	fmt.Printf("The grade are %v\n", grade)
	taskDone := [...]bool{true, true, false, false}
	fmt.Printf("%v\n", taskDone)
	for i := 0; i < len(cities); i++ {
		fmt.Printf("city %v\n", cities[i])
	}
	for index, i := range grade {
		fmt.Printf("index %v value %v\n", index, i)
	}
}
