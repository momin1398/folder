package main

import "fmt"

func main() {
	//code1
	var name string = "xyz"
	country := "xxx"
	fmt.Println(name, country)
	fmt.Printf("He says \"Hello\"")
	fmt.Printf("\nC:\\Users\\a.txt")

	//code2
	r := 'ă'
	fmt.Println(r)
	s1, s2 := "ma", "m"
	s3 := s1 + s2 + string(r)
	fmt.Printf("%v\n", s3)
	//code3
	s4 := "țară means country in Romanian"
	for i := 0; i < len(s4); i++ {
		fmt.Printf("%v\n", s4[i])
	}
	for i, v := range s4 {
		fmt.Printf("The index is %v, the value is %c\n", i, v)
	}
	//code4
	s := "你好 Go!"
	b := []byte(s)
	fmt.Printf("%#v\n", b)
	//iterating over slice
	for i, v := range b {
		fmt.Printf("the index is %v,the value is %v\n", i, v)

	}
	//iterating ovr rune
	for i, v := range b {
		fmt.Printf("the index is %v,the rune is %c\n", i, v)

	}
}
