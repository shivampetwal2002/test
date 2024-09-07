package main

import (
    "fmt"
    "time"
)

// Function that will run as a goroutine

func printMessage(message string) {
    for i := 0; i < 3; i++ {
        fmt.Println(message)
        time.Sleep(time.Second) 
		// Sleep for 1 second
    }
}

func main() {
    // Start two goroutines
    go printMessage("Hello")
    go printMessage("World")

    // Sleep to allow goroutines to finish
    time.Sleep(5 * time.Second)
}

// non deterministic output order 
//output order can vary each time
