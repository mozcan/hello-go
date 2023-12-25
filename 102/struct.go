package main

import (
	"fmt"
	"unsafe"
)

type person struct {
	name    string
	surname string
	age     int16
}

func main() {
	var per person
	fmt.Println(unsafe.Sizeof(per))
	fmt.Println(per)
	person := person{name: "Mustafa", surname: "Özcan", age: 33}
	person.name = "Kemal"
	fmt.Println(person.name)

	anymStruct := struct {
		identity int16
		name     string
	}{12345, "Mustafa"}

	fmt.Println("Identity:", anymStruct.identity, "name:", anymStruct.name)
}
