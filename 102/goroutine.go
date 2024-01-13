package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("This is goroutine Start!!")

	for i := 0; i <= 10; i++ {
		go routine(i)
	}

	fmt.Println("This is goroutine End!!")
	time.Sleep(time.Second)
}

func routine(number int) {
	fmt.Println("This is routine sentence!! ", number)
}
