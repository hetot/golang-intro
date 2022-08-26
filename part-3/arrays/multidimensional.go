package main

import (
	"fmt"
)

func printArray(a [3][2]string) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			fmt.Printf("%s ", a[i][j])
		}
		fmt.Printf("\n")
	}
}

func printArrayUsingRange(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

func main() {
	a := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"}, // This comma is necessary. The compiler will complain if you omit this comma
	}

	printArray(a)

	fmt.Printf("\n")

	var b [3][2]string
	b[0][0] = "apple"
	b[0][1] = "samsung"
	b[1][0] = "microsoft"
	b[1][1] = "google"
	b[2][0] = "AT&T"
	b[2][1] = "T-Mobile"

	printArrayUsingRange(b)
}
