package main

import "fmt"

type dynmVariable interface {
}

func main() {
	var x any

	x = "Mustafa"
	fmt.Println(x)
	x = 33
	fmt.Println(x)

	var dynmVarible dynmVariable
	dynmVarible = "Türkiye"
	fmt.Println("Ülkem:", dynmVarible)
	dynmVarible = 1923
	fmt.Println("Bayram:", dynmVarible)
}
