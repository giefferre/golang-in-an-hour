package main

import (
	"fmt"
	"math/rand"
	"time"
)

func pause() {
	time.Sleep(11 * time.Second)
}

// START1 OMIT
func fuzzBuzz(word string) {
	randomTime := time.Duration(rand.Intn(10000)) * time.Millisecond
	time.Sleep(randomTime)

	fmt.Println(word)
}

func main() {
	// seeding the random generator :)
	rand.Seed(time.Now().UTC().UnixNano())

	go fuzzBuzz("Zak")
	go fuzzBuzz("Yak")
	go fuzzBuzz("Yakketi")
	go fuzzBuzz("Yakk!")

	pause() // waits for a while
	fmt.Println("Quit!")
}

//STOP1 OMIT
