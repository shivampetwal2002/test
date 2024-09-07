package main

import "fmt"

//syntax 
	//type nameOfInterface interface{
		//mtehods
	//}


// Shape interface defination, with an Area method
type Shape interface {
    Area() float64
}

//  struct Rectangle
type Rectangle struct {
    Width  float64
    Height float64
}

// Implement the Area method for Rectangle
func (r Rectangle) Area() float64 {  // method receiver
    return r.Width * r.Height
}

// Square struct
type Square struct {
    Side float64
}

// Implement the Area method for Square
func (s Square) Area() float64 {
    return s.Side * s.Side
}

// Function to print the area of any shape
func printArea(s Shape) {
    fmt.Println("Area:", s.Area())
}

func main() {
    // Create instances of Rectangle and Square
    myRectangle := Rectangle{Width: 5.0, Height: 3.0}
    mySquare := Square{Side: 4.0}
    
    // Print area of both shapes
    printArea(myRectangle)
    printArea(mySquare)
}
