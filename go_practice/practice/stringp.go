package main

import (
	"fmt"
	"strings"
)

func main() {
	var str, str1, str2 string

	fmt.Println("Enter the function to be performed")
	fmt.Scan(&str)
	str = strings.ToLower(str)
	switch str {
	case "upper":
		{

			fmt.Println("Enter the string")
			fmt.Scan(&str1)
			fmt.Printf("changing to upper character %v\n", strings.ToUpper(str1))
		}
	case "lower":
		{

			fmt.Println("Enter the string")
			fmt.Scan(&str1)
			fmt.Printf("changing to upper character %v\n", strings.ToLower(str1))

		}
	case "repeat":
		{

			fmt.Println("Enter string")
			fmt.Scan(&str1)
			fmt.Printf("%v\n", strings.Repeat(str1, 25))
		}
	case "compare":
		{
			fmt.Println("Enter two strings")
			fmt.Scan(&str1, &str2)
			if strings.EqualFold(str1, str2) {
				fmt.Println("thy are equal")
			} else {
				fmt.Println("not equal")
			}

		}

	default:
		fmt.Println("not found")

	}

}
