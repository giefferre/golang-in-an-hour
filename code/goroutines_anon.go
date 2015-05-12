package main

import (
	"fmt"
	"math/rand"
	"time"
)

func pause() {
	time.Sleep(6 * time.Second)
}

// START1 OMIT
func main() {
	rand.Seed(time.Now().UTC().UnixNano()) //OMIT
	go func(word string) {
		randomTime := time.Duration(rand.Intn(5000)) * time.Millisecond
		time.Sleep(randomTime)

		fmt.Println(word)
	}("Zak, Zak, Zak... What am I going to do with you?")

	pause() // waits for a while
	fmt.Println("Quit!")
}

// STOP1 OMIT
