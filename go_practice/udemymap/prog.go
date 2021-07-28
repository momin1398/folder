package main

import "fmt"

func main() {
	//nil map
	//code1
	m1 := make(map[string]bool)
	fmt.Printf("The map m1 %v\n", m1)
	m2 := map[int]string{1: "xyz", 2: "yyy"}
	m2[10] = "zzz"
	fmt.Printf("The map m1 %v\n", m2)
	//code3
	m := map[int]bool{1: true, 2: false, 3: false}
	fmt.Println("before deletng field")
	delete(m, 2)
	fmt.Println("after deletng field", m)
	//m[4] = true
	for index, val := range m {
		fmt.Printf("index %v,value %v\n", index, val)
	}

}
