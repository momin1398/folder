package main

import "fmt"

type perso struct {
	name string
	age  int
	grad grades
}

type grades struct {
	grade  int
	course string
}

func main() {
	me := perso{name: "yyyy", age: 23, grad: grades{
		grade:  89,
		course: "python",
	},
	}
	fmt.Println(me)
	me.grad.grade = 90
	me.grad.course = "golang"
	fmt.Printf("After changes %v\n", me)

}
