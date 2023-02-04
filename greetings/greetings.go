package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Greetings(name string) (string, error) {

	if name == "" {
		return "", errors.New("name can not be empty!")
	}

	message := fmt.Sprintf(randomFormats(), name)
	return message, nil
}

func randomFormats() string {

	formats := []string{
		"Hi, %v. Welcome",
		"Nice to see you , %v",
		"How are you %v ?",
	}

	return formats[rand.Intn(len(formats))]
}
