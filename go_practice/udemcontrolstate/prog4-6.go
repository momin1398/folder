package main

import "fmt"

func main() {
	//code 4
	fmt.Println("The num divisible by 5 & 7 are")
	for i := 1; i <= 500; i++ {
		if i%5 == 0 && i%7 == 0 {
			fmt.Printf("%v\n", i)

		}
	}
	//code5
	fmt.Println("display years from dob to till now")
	for i := 1998; i <= 2021; i++ {
		fmt.Println(i)
	}
	//code 6 using switch
	age := -9
	switch {
	case age < 0 || age > 100:
		fmt.Println("Invalid Age")
	case age < 18:
		fmt.Println("You are minor!")
	case age == 18:
		fmt.Println("Congratulations! You've just become major!")
	default:
		fmt.Println("You are major!")
	}
}
