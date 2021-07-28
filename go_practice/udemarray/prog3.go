package main

import "fmt"

func main() {
	myArray := [3]float64{1.2, 5.6}

	myArray[2] = 6
	fmt.Println(myArray)
	//a := 10 value is int v need to assign float
	a := 10.0
	myArray[0] = a
	fmt.Println(myArray)
	//myArray[3] = 10.10 because this array consist of 3 value since len is 3
	myArray[2] = 10.10
	fmt.Println(myArray)

}
