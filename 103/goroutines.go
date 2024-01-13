package main

import (
	"fmt"
	"time"
)

func count(from string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(from, ":", i)
		time.Sleep(time.Millisecond * 500) // Yarı saniye bekleme
	}
}
func main() {
	// İki farklı fonksiyonu aynı anda çalıştırma (Goroutines)
	go count("Gorutin 1")
	go count("Gorutin 2")
	// Ana fonksiyonun kapanmasını beklemek için bir bekleme süresi ekleyelim
	time.Sleep(time.Second * 3)
}
