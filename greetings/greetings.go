package greetings

import "fmt"

func Greetings(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome.", name)
	return message
}
