package main

import "fmt"

func main() {
	n := 4
	// Read n from input
	fmt.Println(DisplayMinimumNumberFunction(n))
}

// https://www.hackerrank.com/contests/w30/challenges/find-the-minimum-number
func DisplayMinimumNumberFunction(n int) string {
	if n == 2 {
		return "min(int, int)"
	} else {
		return "min(int, " + DisplayMinimumNumberFunction(n-1) + ")"
	}
}
