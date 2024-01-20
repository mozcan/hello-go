package main

import "fmt"

func Print[T any](items []T) {
	for _, v := range items {
		fmt.Println("Item ->", v)
	}
}

func main() {
	intSlice := []int{1, 2, 3, 4, 5, 6}
	Print(intSlice)

	stringSlice := []string{"Mustafa", "Kemal", "Atatürk"}
	Print(stringSlice)

	boolSlice := []bool{true, false, false}
	Print(boolSlice)
}
