package main

import "fmt"

type person struct {
	Name    string
	Age     int
	Country string
}

func main() {
	// Create an instance of person (struct variable (or instance))
	var p person
	//var p2 person
	
	// Initialize the instance
	p.Name = "Shivam"
	p.Age = 22
	p.Country = "India"

	// Print the instance 
	//acessing using DOT operator 
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
	fmt.Println("Country:", p.Country)
}
