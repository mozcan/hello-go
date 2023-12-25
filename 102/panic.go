package main

import "fmt"

func main() {
	process(0)
}

func process(i int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic durumu:", r)
		}
	}()

	if i == 0 {
		panic("Sıfıra Bölünme!!")
	}

	result := 10 / i
	fmt.Println("Sonuç:", result)
}
