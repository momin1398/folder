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

func addEmployee(id int, name string, employeeDetails [3]employee) employee {
	var a employee

	employeeDetails[2].employeeId = id
	employeeDetails[2].employeeName = name
	a = employeeDetails[2]
	return a
}
func employeeData() [3]employee {
	var employeeDetails [3]employee
	employeeDetails[0].employeeId = 11
	employeeDetails[0].employeeName = "abcd"
	employeeDetails[0].employeePhone = 1234
	employeeDetails[1].employeeId = 12
	employeeDetails[1].employeeName = "xyz"
	employeeDetails[1].employeePhone = 4567
	return employeeDetails
}

func main() {

	var employeeDetails [3]employee = employeeData()
	fmt.Println(employeeDetails)
	server()

}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Employee details are: %v", employeeData())
	var b int
	var c string
	fmt.Println("enter the new emp id")
	fmt.Scan(&b)
	fmt.Println("enter the employee")
	fmt.Scan((&c))

	a := addEmployee(b, c, employeeData())
	fmt.Println(a)
	fmt.Fprintf(w, "\nafter adding Employee details are: %v", a)

}
func server() {
	http.HandleFunc("/hellow", ServeHTTP)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
