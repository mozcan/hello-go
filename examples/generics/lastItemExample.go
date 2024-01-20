package main

import "fmt"

func Last[T any](items []T) T {
	return items[len(items)-1]
}

func GenIntLast(items []int) int {
	return items[len(items)-1]
}

func GenStringLast(items []string) string {
	return items[len(items)-1]
}

func main() {

	intSlice := []int{1, 2, 3, 4, 5, 6}
	newSlice := Last[int](intSlice)

	fmt.Println(GenIntLast(intSlice))
	fmt.Println(newSlice)

	stringSlice := []string{"Mustafa", "Kemal", "Atatürk"}
	newStringSlice := Last[string](stringSlice)

	fmt.Println(GenStringLast(stringSlice))
	fmt.Println(newStringSlice)
}
