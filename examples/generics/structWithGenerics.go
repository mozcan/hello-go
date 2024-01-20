package main

import "fmt"

type Bird interface {
	Fly()
	Sound()
}

type Eagle struct {
}

type Owl struct {
}

func (e Eagle) Fly() {
	fmt.Println("Eagle is flying!!")
}

func (e Eagle) Sound() {
	fmt.Println("Eagle is singing!!")
}

func (e Owl) Fly() {
	fmt.Println("Owl is flying!!")
}

func (e Owl) Sound() {
	fmt.Println("Owl is singing!!")
}

func GetStruct[T Bird](bird T) {
	bird.Fly()
	bird.Sound()
}

func main() {
	eagle := Eagle{}
	GetStruct[Bird](eagle)

	owl := Owl{}
	GetStruct[Bird](owl)
}
