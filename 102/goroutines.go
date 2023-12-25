package main

import (
	"fmt"
	"time"
)

func f(arg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(arg, " i :", i)
	}
}

func main() {
	f("Argument")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")

}
