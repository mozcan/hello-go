package main

import "fmt"

func main() {
	count := 2
	messages := make(chan string, count)

	go channelBuffered(messages)
	for i := 1; i < count+1; i++ {
		fmt.Println(i, ".", <-messages)
	}
}

func channelBuffered(messages chan string) {
	messages <- "First Message"
	messages <- "Second Message"
}
