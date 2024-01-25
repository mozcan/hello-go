package main

import (
	"fmt"
	"sync"
)

var result = 0

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go Increment(&wg, &m)
	}

	wg.Wait()
	fmt.Println("Result -> ", result)
}

func Increment(wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()
	result += 1
	mutex.Unlock()

	wg.Done()
}
