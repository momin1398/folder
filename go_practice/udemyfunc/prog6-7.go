package main

import (
	"fmt"
	"strings"
)

func searchItem(b []string, a string) bool {
	for _, i := range b {
		if i == a {
			return true
		}
	}
	return false
}
func searchItemCase(b []string, a string) bool {

	for _, i := range b {
		if strings.EqualFold(i, a) {
			return true
		}
	}
	return false
}
func main() {
	//code 6
	animals := []string{"lion", "tiger", "bear"}
	result := searchItem(animals, "bear")
	fmt.Println(result) // => true
	result = searchItem(animals, "pig")
	fmt.Println(result)
	//code 7
	fmt.Println(strings.Repeat("#", 25))
	animal := []string{"Lion", "tiger", "bear"}
	result = searchItemCase(animal, "beaR")
	fmt.Println(result) // => true
	result = searchItemCase(animal, "lion")
	fmt.Println(result) // => true
}
