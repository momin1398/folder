package main

import "fmt"

func main() {
	var m1 map[string]int
	_ = m1
	marks := make(map[string]int)
	marks["momin"] = 89
	marks["mohd"] = 90
	marks["basha"] = 98
	fmt.Println("the marks are", marks)
	//copying
	frdmarks := marks
	fmt.Println("frd marks", frdmarks)
	frdmarks["momin"] = 99
	fmt.Println("after changing one val in frdmarks then marks will be", marks)
	fmt.Println("copying through iterate")
	m := make(map[string]int)
	for i, val := range marks {
		m[i] = val

	}
	fmt.Println(m)
	m["basha"] = 87
	fmt.Println("after changing val in iterate ", marks)
	delete(m, "basha")
	fmt.Println("after deletng ", m)

}
