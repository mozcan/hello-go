package main

import "fmt"

type human struct {
	name    string
	surname string
	age     int16
}

func (pers human) meth() {
	fmt.Printf("This is a method that works with Struct Person: Name -> %s , Surname -> %s", pers.name, pers.surname)
}

func main() {
	per := human{"Mustafa", "Özcan", 33}
	per.meth()
}
