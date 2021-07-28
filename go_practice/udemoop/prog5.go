package main

import "fmt"

type vehicle interface {
	License() string
	Name() string
}

type car struct {
	licenceNo string
	brand     string
}

func (c car) License() string {
	return c.licenceNo
}
func (c car) Name() string {
	return c.brand
}
func pr(c car) {
	fmt.Println("The details are")
	fmt.Printf("The brand of car is %v\n", c.Name())
	fmt.Printf("The license num is %v\n", c.License())

}

func main() {
	myCar := car{licenceNo: "786546fh", brand: "xyz"}
	pr(myCar)

}
