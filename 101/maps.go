package main

import (
	"fmt"
	"maps"
)

type human struct {
	name, surname string
	age           int16
}

func main() {
	var new_map = make(map[string]int)

	fmt.Println("map:", new_map)
	new_map["Adana"] = 1
	new_map["İstanbul"] = 34
	fmt.Println("new_map", new_map)

	fmt.Println("get:", new_map["Ankara"])
	fmt.Println("len:", len(new_map))

	delete(new_map, "Adana")
	fmt.Println("delete:", new_map)

	clear(new_map)
	fmt.Println("clear:", new_map)

	_, isExist := new_map["Adana"]
	fmt.Println("isExist:", isExist)

	new_map["Adana"] = 1
	value, isExist := new_map["Adana"]
	fmt.Println("isExist:", isExist, "value:", value)

	var equal_map = make(map[string]int)
	equal_map["Adana"] = 1

	if maps.Equal(new_map, equal_map) {
		fmt.Println("new_map == equal_map")
	}

	structMap := make(map[string]human)
	structMap["isim"] = human{"Mustafa", "ÖZCAN", 33}

	for key, val := range structMap {
		fmt.Println("Key:", key, "Value: Name -> ", val.name, "Surname -> ", val.surname, "Age -> ", val.age)
	}
}
