package main

import (
	"errors"
	"fmt"
)

func isNegative(number int) (int, error) {
	if number < 0 {
		return -1, errors.New("This number is negative")
	}

	return number, nil
}

type argError struct {
	arg  int
	prob string
}

func (argStruct *argError) Error() string {
	return fmt.Sprintf("%d - %s", argStruct.arg, argStruct.prob)
}

func isNegativeWithCustomError(number int) (int, *argError) {
	if number < 0 {
		return -1, &argError{number, "can't work with it"}
	}

	return number, nil
}

func main() {
	// Testing error-returning functions with loops.
	for _, i := range []int{7, -4} {
		if r, e := isNegative(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{-1, 42} {
		if r, e := isNegativeWithCustomError(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

}
