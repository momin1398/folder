package main

import (
	"fmt"
)

func main() {
	map1 := make(map[string]int)
	map1["maths"] = 100
	map1["scienc"] = 95
	map1["eng"] = 93
	fmt.Printf("%#v\n", map1)
	fmt.Println("after deleting a sub")
	delete(map1, "eng")
	fmt.Printf("%#v\n", map1) //deleting map
	map1["social"] = 89
	map2 := make(map[string]int)

	for i, v := range map1 { //copying map
		map2[i] = v
	}
	fmt.Printf("The map 1 is%#v\nThe map 2 is %#v\n", map1, map2)
}
