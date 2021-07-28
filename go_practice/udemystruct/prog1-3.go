package main

import "fmt"

type person struct {
	name   string
	age    int
	colors []string
}

func main() {
	//code 1 in udemy
	me := person{name: "xxxx", age: 23, colors: []string{"blue", "red"}}
	you := person{name: "yyyy", age: 24, colors: []string{"pink", "black"}}
	fmt.Println(me)
	fmt.Println(you)
	//code 2
	me.name = "zzzzz"
	var col []string
	col = you.colors
	fmt.Printf("The new variable colors are %v\n", col)
	you.colors = append(you.colors, "white")
	fmt.Printf("The col of you are:%v\n", you.colors)
	//code4
	for q, i := range you.colors {

		fmt.Printf("%v->%v\n", q, i)
	}

}
