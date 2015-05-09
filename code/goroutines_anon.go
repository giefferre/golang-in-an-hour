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
	// seeding the random generator :)
	rand.Seed(time.Now().UTC().UnixNano())

	go func(word string) {
		randomTime := time.Duration(rand.Intn(5000)) * time.Millisecond
		time.Sleep(randomTime)

		fmt.Println(word)
	}("Zack, Zack, Zack... Come devo fare con te?")

	pause()
	fmt.Println("Quit!")
}

// STOP1 OMIT
