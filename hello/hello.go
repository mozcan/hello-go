package main

import (
	"awesomeProject/greetings"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("Hello: ")

	//The flag bits are Ldate, Ltime, and so on.
	log.SetFlags(0)
	//log.SetFlags(1) // Date
	//log.SetFlags(2) // Time

	fmt.Println("Hello World!")

	message, err := greetings.Greetings("")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
