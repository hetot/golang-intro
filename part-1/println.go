package main

import "fmt"

func main() {
	fmt.Println("Binary: %b\\%b", 4, 5)      // Prints `Binary: 100\101`
	fmt.Printf("%d %%\n", 50)                // Prints `50 %`
	s := fmt.Sprintf("Binary: %b\\%b", 4, 5) // s == `Binary: 100\101`
	fmt.Println(s)
	fmt.Println("Binary: %b\\%b", 4) // An argument to Printf is missing.
}
