package main

import "fmt"

func main() {
	one := make(chan string, 1)
	one <- "Message One"
	oneFunc(one)

	two := make(chan string, 1)
	twoFunc(two)
	fmt.Println(<-two)
}

// Receive Only Channel
func oneFunc(one <-chan string) {
	fmt.Println("First message is", <-one)
}

func twoFunc(two chan<- string) {
	two <- "Second message is"
}
