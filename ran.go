package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza shop")
	fmt.Println("Please rate our pizza between 1 to 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// Remove any trailing newline or carriage return characters from the input
	input = strings.TrimSpace(input)

	// Parse the input as a float
	numRating, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Error parsing rating:", err)
		return
	}

	// Adding one to the rating and printing the final result
	fmt.Println("Final rating is:", numRating+1)
} 
