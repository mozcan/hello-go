package main

import "fmt"

func main() {
	var number int = 5
	fmt.Println(plusOne(number))
	var fnumber float64 = 3.1
	fmt.Println(plusOne(fnumber))
}

func plusOne[num int | float64](number num) num {
	return number + 1
}
