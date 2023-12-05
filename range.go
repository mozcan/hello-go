package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4}
	sum := 0

	for _, num := range arr {
		sum += num
	}
	fmt.Println(sum)

	var new_map = make(map[string]int)
	new_map["Adana"] = 1
	new_map["Ankara"] = 6

	for key, value := range new_map {
		fmt.Println("key:", key, "value:", value)
	}

	kvs := map[string]string{"Adana": "kebap"}
	for _, value := range kvs {
		fmt.Println("value:", value)
	}
}
