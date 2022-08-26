// Package declaration
package main

// Factored import statements
import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(math.Max(73.15, 92.46)) // Finding the Max of two numbers

	fmt.Println(math.Sqrt(225)) // Calculate the square root of a number

	fmt.Println(math.Pi) // Printing the value of `𝜋`

	epoch := time.Now().Unix()
	fmt.Println(epoch) // Epoch time in milliseconds

	rand.Seed(epoch)
	fmt.Println(rand.Intn(100)) // Generating a random integer between 0 to 100
}
