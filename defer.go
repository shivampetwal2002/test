package main

import (
    "fmt"
    //"os"
)


// func main() {
//     file, err := os.Open("example.txt")
//     if err != nil {
//         fmt.Println("Error opening file:", err)
//         return
//     }
//     defer file.Close() // Ensures the file is closed when main returns

//     // Perform operations with the file
//     fmt.Println("File opened successfully")
// }



func main() {
    fmt.Println("Start of main function")

    // Call the function that uses defer
    deferDemo()

    fmt.Println("End of main function")
}

func deferDemo() {
    fmt.Println("Start")

    // Defer 
    defer fmt.Println("This is deferred - 1 ")
    defer fmt.Println("This is deferred - 2 ")
    defer fmt.Println("This is deferred - 3 ")

    fmt.Println("End")
}
