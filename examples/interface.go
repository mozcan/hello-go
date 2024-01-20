package main

import "fmt"

type Duck interface {
	Quack()
	DuckGo()
}

type Chicken struct {
}

func (c Chicken) Quack() {
	fmt.Println("Quack")
}

func (c Chicken) DuckGo() {
	fmt.Println("Chicken Go")
}

func DoDuck(d Duck) {
	d.Quack()
	d.DuckGo()
}

func main() {
	c := Chicken{}
	DoDuck(c)
}
