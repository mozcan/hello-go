package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func main() {
	result := plus(1, 2)
	fmt.Println("1+2:", result)

	anymFunc := func(x, y int) int {
		return x + y
	}

	fmt.Println("result:", anymFunc(3, 5))
}
