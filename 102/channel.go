package main

import "fmt"

func main() {
	ch := make(chan string)
	go channel(ch)

	fmt.Println("Channel message is '", <-ch, "'")
}

func channel(ch chan string) {
	ch <- "This is channel message"
	fmt.Println("No receiver! Operation has been blocked")
}
