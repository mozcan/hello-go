package main

import "fmt"

func multipleReturn() (int, int) {
	return 3, 7
}

func main() {
	first, second := multipleReturn()
	fmt.Println("first:", first, "second:", second)

	_, sec := multipleReturn()
	fmt.Println("second:", sec)
}
