package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var lockA, lockB sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			lockA.Lock()
			time.Sleep(100 * time.Millisecond)
			lockB.Lock()
			fmt.Printf("Go Run 1!")
			lockA.Unlock()
			lockB.Unlock()
		}

	}()

	go func() {
		defer wg.Done()
		for {
			lockB.Lock()
			time.Sleep(100 * time.Millisecond)
			lockA.Lock()
			fmt.Printf("Go Run 2!")
			lockB.Unlock()
			lockA.Unlock()
		}
	}()

	wg.Wait()
}
