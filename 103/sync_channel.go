package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	go task(done)

	<-done
}

func task(done chan bool) {
	fmt.Println("Running...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- false
}
