package main

import "fmt"

func main() {

	if 8%4 == 0 {
		fmt.Println("8 is even")
	}

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if num := 7; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 8 {
		fmt.Println(num, "is less than 8")
	} else {
		fmt.Println("Oopps!!")
	}
}
