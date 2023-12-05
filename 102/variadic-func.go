package main

import "fmt"

func sum(nums ...int) int {
	sum := 0
	for _, numb := range nums {
		sum += numb
	}
	fmt.Println(nums)
	return sum
}

func main() {
	numbers := []int{1, 2, 3, 4}
	fmt.Println("numbers:", sum(numbers...))
	fmt.Println("1,2:", sum(1, 2))
	fmt.Println("1,2,3:", sum(1, 2, 3))
}
