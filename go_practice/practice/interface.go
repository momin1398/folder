package main

import (
	"fmt"
	"math"
)

type shape interface {
	volume() float64
	area() float64
}
type sphere struct {
	radius float64
}
type square struct {
	side float64
}

func (c sphere) volume() float64 {
	f := (math.Pi * math.Pow(c.radius, 3)) / 3
	return f * 4

}
func (c sphere) area() float64 {
	f := (math.Pi * math.Pow(c.radius, 2)) * 4
	return f

}
func (a square) area() float64 {
	return math.Pow(a.side, 2)
}
func (a square) volume() float64 {
	fmt.Println("no volume for square")
	return 0
}

func prin(s shape) {
	fmt.Printf("The shape is %#v\n", s)
	fmt.Printf("Area:%v\n", s.area())
	fmt.Printf("volume:%v\n", s.volume())

}

func main() {
	c1 := sphere{4}
	c2 := square{5}
	prin(c1)
	prin(c2)

}
