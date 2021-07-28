package main

import (
	"fmt"
	"net/http"
)

type student struct {
	name string
	age  int
	id   int
}

func studentdetails() [2]student {
	var studentdata [2]student
	studentdata[0].name = "xxxx"
	studentdata[1].name = "xyz"
	studentdata[0].age = 23
	studentdata[1].age = 22
	studentdata[0].id = 9472
	studentdata[1].id = 9421
	return studentdata

}
func findStudent(id int, students [2]student) student {

	var a student
	for i := range students {
		if students[i].id == id {
			a = students[i]
		}
	}
	return a
}

func main() {
	var students [2]student = studentdetails()
	fmt.Println(students)
	serv()
}
func serv() {
	http.HandleFunc("/", ser)
	http.ListenAndServe(":8080", nil)
}
func ser(w http.ResponseWriter, r *http.Request) {
	a := findStudent(9472, studentdetails())
	fmt.Println(a.id)
	fmt.Fprintf(w, "student details are: %v", a)

}
