package main

import "fmt"

func main() {
	var arr [5]int
	fmt.Println("Array is ", arr)

	arr[3] = 100
	fmt.Println("New Array is", arr)
	fmt.Println("get Array element", arr[3])

	fmt.Println("lenght", len(arr))

	newArr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("first element", newArr[0])

	var twoDimen [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoDimen[i][j] = i + j
		}
	}
	fmt.Println("2 Dimensions", twoDimen)
}
