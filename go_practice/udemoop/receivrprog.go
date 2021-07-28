package main

import "fmt"

type student struct {
	name string
	age  int
}

//normal func
func changeStudent(s student, newName string, newAge int) {
	s.name = newName
	s.age = newAge
}

//receiver func
func (s student) changeStudent1(newName string, newAge int) {
	s.name = newName
	s.age = newAge

}
func (s *student) changeStudent2(newName string, newAge int) {
	(*s).name = newName
	(*s).age = newAge

}
func main() {
	st := student{name: "xyx", age: 23}
	fmt.Println("The student details are", st)
	fmt.Println("after calling normal func")
	changeStudent(st, "bbb", 22) //normal method
	fmt.Println(st)
	fmt.Println("after calling receiver func")
	st.changeStudent1("xyz", 24) //calling receiver func
	fmt.Println(st)
	fmt.Println("after calling ptr receiver func")
	st.changeStudent2("zzz", 25)
	fmt.Println(st)
	//copying val
	var s *student
	s = &st
	fmt.Printf("The copy is %v \n", *s)

}
