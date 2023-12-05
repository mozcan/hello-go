package main

import "fmt"

func twoReturnValues() (int, int) {
	return 3, 7
}

func threeReturnValues() (int, string, int) {
	return 1, "Adana", 2
}

func main() {
	first, second := twoReturnValues()
	fmt.Println("first:", first, "second:", second)

	_, sec := twoReturnValues()
	fmt.Println("second:", sec)

	fir, _, third := threeReturnValues()
	fmt.Println("fir:", fir, "third:", third)
}
