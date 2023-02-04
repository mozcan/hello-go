package greetings

import (
	"errors"
	"fmt"
)

func Greetings(name string) (string, error) {

	if name == "" {
		return "", errors.New("name can not be empty!")
	}

	message := fmt.Sprintf("Hi, %v. Welcome.", name)
	return message, nil
}
