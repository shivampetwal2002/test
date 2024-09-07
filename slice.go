package main

import "fmt"

func main() {
    // Creating slices

    // Using make
    s1 := make([]int, 5)
    fmt.Println("s1:", s1) // [0 0 0 0 0]

    // Using slice literals
    s2 := []int{1, 2, 3, 4, 5}
    fmt.Println("s2:", s2) // [1 2 3 4 5]

    // By slicing
    s3 := []int{4, 65, 33, 78, 2, 6}
    subs3 := s3[2:4]
    fmt.Println("subs3:", subs3) // [33 78]

    // Appending
    s4 := []int{3, 3, 3, 3, 3, 3}
    s4 = append(s4, 6, 6, 6, 6)
    fmt.Println("s4 after append:", s4) // [3 3 3 3 3 3 6 6 6 6]

    // Copying
    copy(s1, s4)
    fmt.Println("s1 after copy:", s1) // [3 3 3 3 3]

    // // Removing elements from a slice
    // s5 := []int{1, 2, 3, 4, 5}
    // fmt.Println("s5 before removal:", s5) // [1 2 3 4 5]
    // // Remove element at index 2
    // s5 = append(s5[:2], s5[3:]...)
    // fmt.Println("s5 after removal:", s5) // [1 2 4 5]

    // Using len and cap
    // fmt.Println("Length :", len(s5)) // 4
    // fmt.Println("Capacity :", cap(s5)) // 5 (depends on initial capacity)

    // Iterating over a slice
    fmt.Print("Iterating over s5: ")
    for i, v := range s5 {
        fmt.Printf("Index %d: Value %d ", i, v)
    }
    fmt.Println()

    // Appending with capacity considerations
    s6 := make([]int, 0, 5) // Create slice with capacity of 5
    fmt.Println("Initial capacity of s6:", cap(s6)) // 5
    s6 = append(s6, 1, 2, 3)
    fmt.Println("s6 after initial append:", s6) // [1 2 3]
    fmt.Println("Capacity of s6 after initial append:", cap(s6)) // 5
    s6 = append(s6, 4, 5, 6) // Appending more elements
    fmt.Println("s6 after further append:", s6) // [1 2 3 4 5 6]
    fmt.Println("Capacity of s6 after further append:", cap(s6)) // Capacity may have increased

    // Reslicing
    s7 := []int{10, 20, 30, 40, 50}
    fmt.Println("s7 before reslicing:", s7) // [10 20 30 40 50]
    s7 = s7[1:4]
    fmt.Println("s7 after reslicing:", s7) // [20 30 40]
}
