package main

import "fmt"

type employee struct {
	name    string
	age     int
	sal     float64
	addres  []string
	contact contactdet
}
type contactdet struct {
	phno int
	mail string
}

func main() {
	empdetails := employee{name: "xyz", sal: 600000.0, age: 23, addres: []string{"xxxxxxxx", "yyyyyy"},
		contact: contactdet{phno: 999999, mail: "xxxffh@gmail"}}
	fmt.Printf("The details of emp is%v\n", empdetails)

	//updating phno &age
	empdetails.age = 24
	empdetails.contact.phno = 9999999
	empdetails.addres = append(empdetails.addres, "zzzzzz")
	fmt.Printf("The updated details %v\n", empdetails)

}
