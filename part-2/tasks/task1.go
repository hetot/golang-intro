package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		FizzBuzz(i)
	}
}

func FizzBuzz(x int) {
	ans := ""
	if x%3 == 0 {
		ans += "Fizz"
	}
	if x%5 == 0 {
		ans += "Buzz"
	}
	fmt.Printf("%d - %s\n", x, ans)
}
