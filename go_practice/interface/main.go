package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}
type circle struct {
	radius float64
}
type rectangle struct {
	height, width float64
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func (r rectangle) perimeter() float64 {
	return r.height * r.width
}

func (c circle) area() float64 { //receiver func
	return math.Pi * math.Pow(c.radius, 2)
}
func (r rectangle) area() float64 {
	return r.height * r.width
}
func prin(s shape) {
	fmt.Printf("The shape is %#v\n", s)
	fmt.Printf("Area:%v\n", s.area())
	fmt.Printf("perimeter:%v\n", s.perimeter())
}
func main() {
	c1 := circle{6}
	c2 := rectangle{10, 3}
	prin(c1)
	prin(c2)

}
