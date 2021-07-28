package main

import (
	"fmt"
	"strings"
)

var stud = map[string]int{}

func enterSub(n int) []string {

	var m string
	var sub []string
	fmt.Printf("Enter subject name\n")
	for i := len(sub); i < n; i++ {
		fmt.Scan(&m)
		sub = append(sub, m)

	}
	return sub
}
func addsub(n int, l int, sub ...string) {
	var m int

	for i := l; i < n; i++ {
		fmt.Printf("Enter the mark of %v ", sub[i])
		fmt.Scan(&m)
		stud[sub[i]] = m

	}

}
func find(n string) (int, string) {
	for index, value := range stud {
		if index == n {
			return value, index

		}

	}
	return 0, "null"
}
func main() {

	var n, l int
	var fu string
	l = 0
	fmt.Println("Enter the no of subjects")
	fmt.Scan(&n)
	a := enterSub(n)
	fmt.Println("Enter marks")

	addsub(n, l, a...)
	fmt.Printf("\nThe marks are %v\n", stud)
	fmt.Println("Enter function to be perform")
	fmt.Scan(&fu)
	fu = strings.ToLower(fu)
	switch fu {
	case "add":
		{
			var s string
			fmt.Printf("enter new sub\n")
			fmt.Scan(&s)
			a = append(a, s)
			l = len(a)
			addsub(n+1, l-1, a...)

			fmt.Printf("\nThe marks are %v\n", stud)

		}
	case "find":
		{
			var s string
			fmt.Println("enter sub")
			fmt.Scan(&s)
			n, _ := find(s)
			if n != 0 {
				fmt.Printf("The mark of %v is %v\n", s, n)
			} else {
				fmt.Println("sub not found")
			}
		}
	case "delete":
		{
			var s string
			fmt.Println("enter sub to be removed")
			fmt.Scan(&s)
			_, sub := find(s)
			delete(stud, sub)
			fmt.Printf("The current student details %v\n", stud)

		}
	case "update":
		{
			var s string
			var m int
			fmt.Println("enter sub to be searched")
			fmt.Scan(&s)
			fmt.Printf("correct marks of %v\n", s)
			fmt.Scan(&m)
			_, n := find(s)
			stud[n] = m
			fmt.Printf("the updated marks are %v\n", stud)
		}
	default:
		fmt.Println("not found")

	}

}
