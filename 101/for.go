package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for j := 5; j <= 10; j++ {
		fmt.Println(j)
	}

	for k := 0; k <= 10; k++ {
		fmt.Println(k)
		if k == 3 {
			fmt.Println("Break")
			break
		}
	}

	for n := 0; n <= 10; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
