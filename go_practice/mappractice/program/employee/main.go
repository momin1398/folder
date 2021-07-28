package main

import (
	"fmt"
	"strings"
)

type employee struct {
	name   string
	age    int
	salary int
}

var emplo = map[employee]int{}
var empl []employee

func add(n int, l int) {
	var emp employee
	for i := l; i < n; i++ {
		fmt.Printf("enter name\n")
		fmt.Scan(&emp.name)
		fmt.Printf("enter age of %v\n", emp.name)
		fmt.Scan(&emp.age)
		fmt.Printf("Enter sal of %v\n", emp.name)
		fmt.Scan(&emp.salary)
		empl = append(empl, emp)

	}

}

func per(n int, l int) {
	var id int

	for i := l; i < n; i++ {
		fmt.Printf("enter id of %v\n", empl[i].name)
		fmt.Scan(&id)
		emplo[empl[i]] = id
	}
	fmt.Printf("The employee details are %v\n", emplo)

}
func find(name string) employee {

	for index, _ := range empl {
		if empl[index].name == name {
			return empl[index]

		}
	}
	return employee{}

}
func swi(fu string, n int) {
	fu = strings.ToLower(fu)

	switch fu {
	case "add":
		{
			add(n+1, n)
			per(n+1, n)
		}
	case "find":
		{
			var name string
			fmt.Println("enter the name to be find")
			fmt.Scan(&name)
			index := find(name)
			if index != (employee{}) {
				fmt.Printf("The name founded %v\n at the index", emplo[index])
			} else {
				fmt.Println("not found")
			}

		}
	case "delete":
		{
			var name string
			fmt.Println("enter the name to be deleted")
			fmt.Scan(&name)
			index := find(name)
			delete(emplo, index)
			fmt.Printf("The current employee are %v\n", emplo)
		}
	case "update":
		{
			var name string
			var n int
			fmt.Println("enter the name to be updated")
			fmt.Scan(&name)
			fmt.Printf("enter new id of %v\n", name)
			fmt.Scan(&n)
			index := find(name)
			emplo[index] = n
			fmt.Printf("The current employee are %v\n", emplo)

		}
	default:
		fmt.Println("not found")
	}

}

func main() {
	var n int
	var fu string
	l := 0
	fmt.Println("enter the number of employee")
	fmt.Scan(&n)
	add(n, l)
	per(n, l)
	fmt.Println("enter func to be performed")
	fmt.Scan(&fu)
	swi(fu, n)

}
