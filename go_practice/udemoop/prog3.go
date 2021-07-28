package main

import "fmt"

type boo struct {
	price float64
	book  string
}

func (b boo) vat() float64 {
	return b.price * 0.09
}
func (b *boo) discount() {
	(*b).price = b.price * 0.9
}

func main() {
	b := boo{100, "xyz"}
	v := b.vat()
	fmt.Println("the vat is ", v)
	b.discount()
	fmt.Printf("The updated price is %v\n", b.price)
}
