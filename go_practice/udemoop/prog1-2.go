package main

import "fmt"

type money float64

func (m money) printt() {
	fmt.Printf("in 2 decimal points %.2f\n", m)
}
func (m money) printstr() string {
	return fmt.Sprintf("%.2f", m)

}

func main() {
	var mon money = 56.0
	mon.printt()
	//code 2
	s := mon.printstr()
	fmt.Printf("the float to string %#v\n ", s)

}
