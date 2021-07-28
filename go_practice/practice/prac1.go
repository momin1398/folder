package main

import (
	"fmt"
	"math"
	"strings"
)

type shap interface {
	area() float64
	perimeter() float64
}
type circle struct {
	radius float64
}
type rectangle struct {
	width, length float64
}
type squar struct {
	side float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
func (c circle) perimeter() float64 {
	return math.Pi * 2 * c.radius
}
func (r rectangle) area() float64 {
	return r.length * r.width
}
func (r rectangle) perimeter() float64 {
	return 2 * (r.length + r.width)
}
func (a squar) area() float64 {
	return a.side * a.side
}
func (a squar) perimeter() float64 {
	return 4 * a.side

}
func prinf(s shap) {
	fmt.Printf("\nThe shape is %T\n", s)
	fmt.Printf("The area of %T is %v\n", s, s.area())
	fmt.Printf("The perimeter of %T is %v\n", s, s.perimeter())
}

func main() {
	fmt.Println("enter shape")
	var s string
	fmt.Scan(&s)
	s = strings.ToLower(s)
	switch s {
	case "circle":
		{
			var r float64
			fmt.Println("Enter the radius")
			fmt.Scan(&r)
			c1 := circle{radius: r}
			prinf(c1)

		}
	case "rectangle":
		{
			var l, w float64
			fmt.Println("Enter the length & width")
			fmt.Scan(&l, &w)
			r := rectangle{length: l, width: w}
			prinf(r)

		}
	case "square":
		{
			var a float64
			fmt.Println("Enter side")
			fmt.Scan(&a)
			sq := squar{side: a}
			prinf(sq)

		}
	default:
		fmt.Println("shape not found")
	}

}
