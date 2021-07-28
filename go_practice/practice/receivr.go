package main

import "fmt"

type emp struct {
	name string
	age  int
	sal  int
}

//normal func
func changedetails(e emp, n string, a int, s int) {
	e.name = n
	e.sal = s
	e.age = a
}
func (e emp) changedet(n string, a int, s int) {
	e.name = n
	e.sal = s
	e.age = a

}
func (e *emp) changedetptr(n string, a int, s int) {
	(*e).name = n
	(*e).sal = s
	(*e).age = a
}
func main() {
	e1 := emp{name: "xyz", age: 23, sal: 300000}
	fmt.Println("initial details are", e1)
	changedetails(e1, "xxx", 24, 679000)
	fmt.Println("calling normal func and result is", e1)
	e1.changedet("yyy", 22, 768766)
	fmt.Println("calling normal receiver func and result is", e1)
	e1.changedetptr("mmmm", 22, 768766)
	fmt.Println("calling pointr receiver func and result is", e1)

}
