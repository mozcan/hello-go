package main

import (
	"fmt"
	"math"
)

const constant string = "constant"

func main() {
	fmt.Println(constant)

	const number = 5000000
	const result = 3e20 / number
	fmt.Println(result)
	fmt.Println(int64(result))
	fmt.Println(math.Sin(result))
}
