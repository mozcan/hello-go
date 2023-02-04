package main

import (
	"awesomeProject/greetings"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("Hello: ")
	log.SetFlags(0)

	fmt.Println("Hello World!")

	message, err := greetings.Greetings("")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
