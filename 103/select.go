package main

import (
	"fmt"
	"time"
)

func main() {
	number := make(chan int)
	message := make(chan string)

	go channelOne(number)
	go channelTwo(message)

	select {
	case firstChannel := <-number:
		fmt.Println("First Channel message", firstChannel)

	case secondChannel := <-message:
		fmt.Println("Second Channel message", secondChannel)

		/*
			default:
				fmt.Println("Wait!! Channels are not ready for execution")

		*/
	}
}

func channelOne(number chan int) {
	time.Sleep(2 * time.Second)
	number <- 33
}

func channelTwo(msg chan string) {
	time.Sleep(2 * time.Second)
	msg <- "This is message"
}
