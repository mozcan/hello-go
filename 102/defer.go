package main

import "fmt"

func main() {
	anymFunc := func(x, y int) int {
		return x * y
	}

	defer fmt.Println("First Chapter")
	defer fmt.Println("Anonyms Func Result : ", anymFunc(4, 6))
	fmt.Println("Second Chapter")
	defer fmt.Println("Third Chapter")
	fmt.Println("Last Chapter")
}
