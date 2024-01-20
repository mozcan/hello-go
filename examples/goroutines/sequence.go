package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	evenCh, oddCh := make(chan bool, 1), make(chan bool, 1)
	defer close(evenCh)
	defer close(oddCh)

	wg = sync.WaitGroup{}
	wg.Add(2)

	go even(evenCh, oddCh)
	go odd(oddCh, evenCh)

	// initiate the odd function to start first
	evenCh <- true

	wg.Wait()
}

func even(evenCh, oddCh chan bool) {
	for i := 2; i <= 5; i += 2 {
		<-oddCh
		fmt.Println("Even ->", i)
		evenCh <- true
	}

	wg.Done()
}

func odd(oddCh, evenCh chan bool) {
	for i := 1; i <= 5; i += 2 {
		<-evenCh
		fmt.Println("Odd ->", i)
		oddCh <- true
	}

	wg.Done()
}
