package main

import (
	"fmt"
)

func main() {
	//code1
	a1 := []int{3, 6, 8}
	fmt.Println("iterate over the slice")
	for i, val := range a1 {
		fmt.Printf("The index %v, value is %v\n", i, val)
	}
	//code 3
	num := []float64{7.9, 9.4, 9.8}
	num = append(num, 10.1)
	num = append(num, 4.1, 5.5, 6.1)
	fmt.Println(num)
	n := []float64{8.9, 7.0}
	num = append(num, n...)
	for i, val := range num {
		fmt.Printf("The index %v, value is %v\n", i, val)
	}
	//code5
	fmt.Println("ignoring first and last 2 num")
	nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}
	s := 0
	for i, v := range nums[2 : len(nums)-2] {
		fmt.Printf("The index %v, value is %v\n", i, v)
		s += v

	}
	fmt.Println("The sum is ", s)
	//code7
	friends := []string{"Mmmm", "yyy", "zzz", "xyz"}
	var myfriend []string
	myfriend = append(myfriend, friends...)
	myfriend[2] = "harry"
	fmt.Println(friends, myfriend)

	//code8
	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
	var myyear []int
	myyear = append(years[:3], years[len(years)-3:]...)
	fmt.Println(myyear)

}
