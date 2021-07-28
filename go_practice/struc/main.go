package main

import "fmt"

type contact struct {
	phno   int
	mailid string
}
type student struct {
	name        string
	age         int
	contactinfo contact
}

func main() {
	xyz := student{
		name: "xyz xxx yyy",
		age:  23,
		contactinfo: contact{
			phno:   74568990,
			mailid: "mxyz@gmail.com",
		},
	}
	fmt.Printf("The details of student is %#v \n", xyz)
}
