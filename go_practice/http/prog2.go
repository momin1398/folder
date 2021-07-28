package main

import (
	"fmt"
	"log"
	"net/http"
)

type employee struct {
	employeeId    int
	employeeName  string
	employeePhone int
}

func findEmployee(id int, employeeDetails [2]employee) employee {
	var a employee
	for i := range employeeDetails {
		if employeeDetails[i].employeeId == id {
			a = employeeDetails[i]
		}
	}
	return a
}
func employeeData() [2]employee {
	var employeeDetails [2]employee
	employeeDetails[0].employeeId = 11
	employeeDetails[0].employeeName = "xyz"
	employeeDetails[0].employeePhone = 1234
	employeeDetails[1].employeeId = 12
	employeeDetails[1].employeeName = "abc"
	employeeDetails[1].employeePhone = 4567
	return employeeDetails
}

func main() {

	var employeeDetails [2]employee = employeeData()
	fmt.Println(employeeDetails)
	server()

}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {

	a := findEmployee(11, employeeData())
	fmt.Println(a.employeeId)
	fmt.Fprintf(w, "Employee details are: %v", a)

}
func server() {
	http.HandleFunc("/hellow", ServeHTTP)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
