package main

import (
	"fmt"
	"math/rand"
	"time"
)

func pause() {
	time.Sleep(20 * time.Second)
}

// START1 OMIT
func beacon(ch chan string) {
	for {
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		ch <- "Ping"
	}
}

func aMoreSecureShip(ch chan string) {
	ticker := time.NewTicker(750 * time.Millisecond)

	for {
		select {
		case <-ticker.C:
			fmt.Println("No signal from beacon, standing by...")
		case data := <-ch:
			fmt.Println("Received", data)
		}
	}
}

// STOP1 OMIT

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	channel := make(chan string)

	go beacon(channel)
	go aMoreSecureShip(channel)

	pause()
	fmt.Println("Quit!")
}
