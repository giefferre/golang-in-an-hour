package main

import (
    "fmt"
)

func main() {
    // START1 OMIT
    aValue := "Change me!"

    if true {
        aValue := "!em egnahC" // Something's wrong here

        fmt.Println(aValue)
    }

    fmt.Println(aValue)
    // STOP1 OMIT
}