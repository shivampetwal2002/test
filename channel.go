// Channels are used to communicate between goroutines and synchronize
// their execution by allowing them to send and receive data.

// Channels are created using the make function.

// ch := make(chan int) // Creates a channel for int data
// ch <- 42 // Sends the value 42 into the channel
// value := <-ch // Receives a value from the channel


// basic 

package main

import (
    "fmt"
)

func main() {
    ch := make(chan string) // Create a new channel

    // Start a goroutine that sends data into the channel
    go func() {
        ch <- "mesaage from goroutine! using channel" // Send a message to channel
    }()

// Receive data from the channel
    message := <-ch
    fmt.Println(message) 
}
