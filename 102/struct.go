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
	fmt.Println(person.name)
}
